package cwaqr

import (
	"fmt"

	"github.com/lukasmalkmus/cwa-qr/internal/pb"
)

//go:generate bin/stringer -type=TemporaryEventType,PermanentEventType -linecomment -output=type_string.go

// EventType is the type of an event.
type EventType interface {
	fmt.Stringer

	// Make sure implementations of EventType can only originate in this
	// package.
	et()
}

func (et TemporaryEventType) et() {}
func (et PermanentEventType) et() {}

// TemporaryEventType describes the type of a temporary event.
type TemporaryEventType uint8

// All available temporary event types.
const (
	OtherTemporary TemporaryEventType = iota + 1 // Other Temporary
	CulturalEvent                                // Cultural Event
	ClubActivity                                 // Club Activity
	PrivateEvent                                 // Private Event
	WorshipService                               // Worship service
)

// PermanentEventType describes the type of a permanent event.
type PermanentEventType uint8

// All available permanent event types.
const (
	OtherPermanent         PermanentEventType = iota + 1 // Other Permanent
	Retail                                               // Retail
	FoodService                                          // Food Service
	Craft                                                // Craft
	Workplace                                            // Workplace
	EducationalInstitution                               // Educational Institution
	PublicBuilding                                       // Public Building
)

var pbEventTypes = map[EventType]pb.TraceLocationType{
	// Temporary
	OtherTemporary: pb.TraceLocationType_LOCATION_TYPE_TEMPORARY_OTHER,
	CulturalEvent:  pb.TraceLocationType_LOCATION_TYPE_TEMPORARY_CULTURAL_EVENT,
	ClubActivity:   pb.TraceLocationType_LOCATION_TYPE_TEMPORARY_CLUB_ACTIVITY,
	PrivateEvent:   pb.TraceLocationType_LOCATION_TYPE_TEMPORARY_PRIVATE_EVENT,
	WorshipService: pb.TraceLocationType_LOCATION_TYPE_TEMPORARY_WORSHIP_SERVICE,

	// Permanent
	OtherPermanent:         pb.TraceLocationType_LOCATION_TYPE_PERMANENT_OTHER,
	Retail:                 pb.TraceLocationType_LOCATION_TYPE_PERMANENT_RETAIL,
	FoodService:            pb.TraceLocationType_LOCATION_TYPE_PERMANENT_FOOD_SERVICE,
	Craft:                  pb.TraceLocationType_LOCATION_TYPE_PERMANENT_CRAFT,
	Workplace:              pb.TraceLocationType_LOCATION_TYPE_PERMANENT_WORKPLACE,
	EducationalInstitution: pb.TraceLocationType_LOCATION_TYPE_PERMANENT_EDUCATIONAL_INSTITUTION,
	PublicBuilding:         pb.TraceLocationType_LOCATION_TYPE_PERMANENT_PUBLIC_BUILDING,
}

func getProtoEventType(eventType EventType) (pb.TraceLocationType, error) {
	pbEventType, ok := pbEventTypes[eventType]
	if !ok {
		return 0, fmt.Errorf("unknown event type %q", eventType.String())
	}
	return pbEventType, nil
}
