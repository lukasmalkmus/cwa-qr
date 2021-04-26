package cwaqr

import (
	"fmt"
	"time"
)

// CharacterLimit is the character limit for user provided strings.
const CharacterLimit = 100

// Event which users of the Corona Warn App check into.
type Event struct {
	// Description of the event. Limited to 100 characters.
	Description string
	// Address of the event. Limited to 100 characters.
	Address string
	// Type of the event. There are types for permanent and temporary events.
	Type EventType
	// Duration is the average duration a person stays at an event. Only valid
	// for temporary events.
	Duration time.Duration
	// Start of the event. Only valid for permanent events.
	Start time.Time
	// End of the event. Only valid for permanent events.
	End time.Time
}

// Validate the event.
func (e *Event) Validate() error {
	if l := len(e.Description); l > CharacterLimit {
		return fmt.Errorf("description exceeds character limit of %d by %d", CharacterLimit, l-CharacterLimit)
	}
	if l := len(e.Address); l > CharacterLimit {
		return fmt.Errorf("description exceeds character limit of %d by %d", CharacterLimit, l-CharacterLimit)
	}

	switch et := e.Type.(type) {
	case TemporaryEventType:
		if e.Duration == 0 {
			return fmt.Errorf("temporary event %q must have duration set", et)
		}
		if !e.Start.IsZero() {
			return fmt.Errorf("temporary event %q cannot have start time set", et)
		}
		if !e.End.IsZero() {
			return fmt.Errorf("temporary event %q cannot have end time set", et)
		}
	case PermanentEventType:
		if e.Duration != 0 {
			return fmt.Errorf("permanent event %q cannot have duration set", et)
		}
		if e.Start.IsZero() {
			return fmt.Errorf("permanent event %q must have start time set", et)
		}
		if e.End.IsZero() {
			return fmt.Errorf("permanent event %q must have end time set", et)
		}
	default:
		return fmt.Errorf("invalid event type %q", et)
	}

	return nil
}

// NewTemporaryEvent creates a new temporary event with the given type and
// average stay duration.
func NewTemporaryEvent(description, address string, eventType TemporaryEventType, duration time.Duration) Event {
	return Event{
		Description: description,
		Address:     address,
		Type:        eventType,
		Duration:    duration,
	}
}

// NewPermanentEvent creates a new permanent event with the given type and
// stay interval.
func NewPermanentEvent(description, address string, eventType PermanentEventType, start, end time.Time) Event {
	return Event{
		Description: description,
		Address:     address,
		Type:        eventType,
		Start:       start,
		End:         end,
	}
}
