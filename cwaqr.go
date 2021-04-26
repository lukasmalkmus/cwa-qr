package cwaqr

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/skip2/go-qrcode"
	"google.golang.org/protobuf/proto"

	"github.com/lukasmalkmus/cwa-qr/internal/pb"
)

const (
	baseUrl   = "https://e.coronawarn.app?v=1#"
	publicKey = "gwLMzE153tQwAOf2MZoUXXfzWTdlSpfS99iZffmcmxOG9njSK4RTimFOFwDh6t0Tyw8XR01ugDYjtuKwjjuK49Oh83FWct6XpefPi9Skjxvvz53i9gaMmUEc96pbtoaA"
)

// GenerateQRCode creates a new QR-Code for the given event which can be scanned
// by the Corona Warn App. The code is 512x512 in size with a recovery level of
// "M" (15%).
func GenerateQRCode(event Event) ([]byte, error) {
	qrCodeContent, err := GenerateURLString(event)
	if err != nil {
		return nil, err
	}
	return qrcode.Encode(qrCodeContent, qrcode.Medium, 512)
}

// GenerateURLString creates a new URL string which represents the content of a
// QR-Code.
func GenerateURLString(event Event) (string, error) {
	if err := event.Validate(); err != nil {
		return "", err
	}

	// Create the base payload.
	qrCodePayload := &pb.QRCodePayload{
		Version: 1,
	}

	// Attach trace location.
	qrCodePayload.LocationData = &pb.TraceLocation{
		Version:     1,
		Description: event.Description,
		Address:     event.Address,
	}

	// Attach crowd notifier data.
	var err error
	if qrCodePayload.CrowdNotifierData, err = createCrowdNotifierData(); err != nil {
		return "", err
	}

	// Create vendor data.
	vendorData, err := createVendorData(event.Type)
	if err != nil {
		return "", err
	}

	// Depending on the event type, add the average stay duration or stay
	// interval to the payload.
	switch et := event.Type.(type) {
	case TemporaryEventType:
		eventDuration := event.Duration.Truncate(time.Minute)
		vendorData.DefaultCheckInLengthInMinutes = uint32(eventDuration.Minutes())
	case PermanentEventType:
		qrCodePayload.LocationData.StartTimestamp = uint64(event.Start.Unix())
		qrCodePayload.LocationData.EndTimestamp = uint64(event.End.Unix())
	default:
		return "", fmt.Errorf("invalid event type %q", et)
	}

	// Attach vendor data.
	if qrCodePayload.VendorData, err = proto.Marshal(vendorData); err != nil {
		return "", err
	}

	// Encode the payload to its protobuf representation.
	qrCodePb, err := proto.Marshal(qrCodePayload)
	if err != nil {
		return "", err
	}

	// Base64 encode the QR-Code protobuf message and generate the payload URL.
	return baseUrl + base64.URLEncoding.EncodeToString(qrCodePb), nil
}

func createCrowdNotifierData() (*pb.CrowdNotifierData, error) {
	rawPublicKey, err := base64.RawStdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}

	cryptoSeed := make([]byte, 16)
	if _, err = rand.Read(cryptoSeed); err != nil {
		return nil, err
	}

	return &pb.CrowdNotifierData{
		Version:           1,
		PublicKey:         rawPublicKey,
		CryptographicSeed: cryptoSeed,
	}, nil
}

func createVendorData(eventType EventType) (*pb.CWALocationData, error) {
	pbEventType, err := getProtoEventType(eventType)
	if err != nil {
		return nil, err
	}

	return &pb.CWALocationData{
		Version: 1,
		Type:    pbEventType,
	}, nil
}
