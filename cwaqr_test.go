package cwaqr_test

import (
	"testing"
	"time"

	cwaqr "github.com/lukasmalkmus/cwa-qr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate_PermanentEvent(t *testing.T) {
	b, err := cwaqr.Generate(cwaqr.Event{
		Description: "Work, Inc.",
		Address:     "Work City",
		Type:        cwaqr.Workplace,
		Duration:    time.Hour * 12,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, b)
}

func TestGenerate_TemporaryEvent(t *testing.T) {
	b, err := cwaqr.Generate(cwaqr.Event{
		Description: "Party",
		Address:     "Home",
		Type:        cwaqr.PrivateEvent,
		Duration:    time.Hour * 4,
		Start:       time.Now(),
		End:         time.Now().Add(time.Hour * 4),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, b)
}
