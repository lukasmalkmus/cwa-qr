package cwaqr

import (
	"errors"
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
	// Type of the event which is some kind of either permanent or temporary
	// event.
	Type EventType
	// Duration is the average duration a person stays at an event.
	Duration time.Duration
	// Start of the event. Only valid for temporary events.
	Start time.Time
	// End of the event. Only valid for temporary events.
	End time.Time
}

// Validate the event.
func (e *Event) Validate() error {
	if l := len(e.Description); l > CharacterLimit {
		return fmt.Errorf("description exceeds character limit of %d by %d bytes", CharacterLimit, l-CharacterLimit)
	}
	if l := len(e.Address); l > CharacterLimit {
		return fmt.Errorf("address exceeds character limit of %d by %d bytes", CharacterLimit, l-CharacterLimit)
	}

	if e.Duration == 0 {
		return errors.New("event must have duration set")
	}

	if e.Type.IsTemporary() {
		if e.Start.IsZero() {
			return fmt.Errorf("temporary event %q must have start time set", e.Type)
		}
		if e.End.IsZero() {
			return fmt.Errorf("temporary event %q must have end time set", e.Type)
		}
	}

	if !e.Type.IsPermanent() && !e.Type.IsTemporary() {
		return fmt.Errorf("invalid event %q", e.Type)
	}

	return nil
}
