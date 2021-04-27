package cwaqr_test

import (
	"testing"
	"time"

	cwaqr "github.com/lukasmalkmus/cwa-qr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateURLString_PermanentEvent(t *testing.T) {
	s, err := cwaqr.GenerateURLString(cwaqr.Event{
		Description: "Work, Inc.",
		Address:     "Work City",
		Type:        cwaqr.Workplace,
		Duration:    time.Hour * 12,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, s)
}

func TestGenerateURLString_TemporaryEvent(t *testing.T) {
	s, err := cwaqr.GenerateURLString(cwaqr.Event{
		Description: "Party",
		Address:     "Home",
		Type:        cwaqr.PrivateEvent,
		Duration:    time.Hour * 4,
		Start:       time.Now(),
		End:         time.Now().Add(time.Hour * 4),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, s)
}
