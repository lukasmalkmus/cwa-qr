package cwaqr

import (
	"crypto/rand"
	"encoding/base64"
	"net/url"
	"time"

	"github.com/skip2/go-qrcode"
	"google.golang.org/protobuf/proto"

	"github.com/lukasmalkmus/cwa-qr/internal/pb"
)

const (
	baseUrl   = "https://e.coronawarn.app?v=1#"
	publicKey = "gwLMzE153tQwAOf2MZoUXXfzWTdlSpfS99iZffmcmxOG9njSK4RTimFOFwDh6t0Tyw8XR01ugDYjtuKwjjuK49Oh83FWct6XpefPi9Skjxvvz53i9gaMmUEc96pbtoaA"
)

// Generate creates a new QR-Code for the given event which can be scanned by
// the Corona Warn App. The code is 512x512 in size with a recovery level of "L"
// (low = 7%).
func Generate(event Event) ([]byte, error) {
	qrCodeURL, err := GenerateURL(event)
	if err != nil {
		return nil, err
	}
	return qrcode.Encode(qrCodeURL.String(), qrcode.Low, 512)
}

// GenerateURL creates a new URL string which represents the content of a
// QR-Code.
func GenerateURL(event Event) (*url.URL, error) {
	if err := event.Validate(); err != nil {
		return nil, err
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
		return nil, err
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
		return nil, err
	}

	// Base64 encode the QR-Code protobuf message and construct the payload URL.
	qrCodeStr := base64.URLEncoding.EncodeToString(qrCodePb)
	u, err := url.ParseRequestURI(baseUrl + qrCodeStr)
	if err != nil {
		return nil, err
	}

	return u, nil
}
