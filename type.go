package cwaqr

import "github.com/lukasmalkmus/cwa-qr/internal/pb"

//go:generate bin/stringer -type=EventType -linecomment -output=type_string.go

// EventType describes the type of an event. Types belong to one of two
// categories: permanent or temporary.
type EventType uint8

// IsPermanent return true if the event type describes a permanent event.
func (et EventType) IsPermanent() bool {
	return et == OtherPermanent || (et >= Retail && et <= PublicBuilding)
}

// IsTemporary return true if the event type describes a temporary event.
func (et EventType) IsTemporary() bool {
	return et == OtherTemporary || (et >= CulturalEvent && et <= WorshipService)
}

func (et EventType) toProto() pb.TraceLocationType {
	if et.IsPermanent() || et.IsTemporary() {
		return pb.TraceLocationType(et)
	}
	return pb.TraceLocationType_LOCATION_TYPE_UNSPECIFIED
}

// All available  event types.
const (
	Unknown EventType = iota // Unknown

	OtherPermanent // Other Permanent
	OtherTemporary // Other Temporary

	Retail                 // Retail
	FoodService            // FoodS ervice
	Craft                  // Craft
	Workplace              // Workplace
	EducationalInstitution // Educational Institution
	PublicBuilding         // Public Building

	CulturalEvent  // Cultural Event
	ClubActivity   // Club Activity
	PrivateEvent   // Private Event
	WorshipService // Worship service
)
