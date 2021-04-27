package cwaqr

import (
	"crypto/rand"
	"encoding/base64"
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

	location := &pb.TraceLocation{
		Version:     1,
		Description: event.Description,
		Address:     event.Address,
	}

	if event.Type.IsTemporary() {
		location.StartTimestamp = uint64(event.Start.Unix())
		location.EndTimestamp = uint64(event.End.Unix())
	}

	cryptoSeed := make([]byte, 16)
	if _, err := rand.Read(cryptoSeed); err != nil {
		return "", err
	}

	crowdNotifier := &pb.CrowdNotifier{
		Version:           1,
		PublicKey:         []byte(publicKey),
		CryptographicSeed: cryptoSeed,
	}

	eventDuration := event.Duration.Truncate(time.Minute)
	vendor := &pb.Vendor{
		Version:                       1,
		Type:                          event.Type.toProto(),
		DefaultCheckInLengthInMinutes: uint32(eventDuration.Minutes()),
	}

	// Create the base payload.
	qrCodePayload := &pb.QRCodePayload{
		Version:       1,
		Location:      location,
		CrowdNotifier: crowdNotifier,
		Vendor:        vendor,
	}

	// Encode the payload to its protobuf representation.
	qrCodePb, err := proto.Marshal(qrCodePayload)
	if err != nil {
		return "", err
	}

	// Base64 encode the QR-Code protobuf message and generate the payload URL.
	return baseUrl + base64.URLEncoding.EncodeToString(qrCodePb), nil
}
