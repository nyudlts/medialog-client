package medialog

import (
	"flag"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	config      string
	environment string
)

func init() {
	flag.StringVar(&config, "config", "", "")
	flag.StringVar(&environment, "environment", "", "")
}

func TestClient(t *testing.T) {
	flag.Parse()
	var mlc *MedialogClient
	t.Run("test create client", func(t *testing.T) {

		var err error
		mlc, err = NewClient(20)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("test get medialog info", func(t *testing.T) {

		mlInfo, err := mlc.GetHostInfo()
		if err != nil {
			t.Error(err)
		}

		t.Log(mlInfo)
	})

	var entryUUID uuid.UUID
	t.Run("test get all medialog entry uuids", func(t *testing.T) {
		entries, err := mlc.GetEntryUUIDs()
		if err != nil {
			t.Error(err)
		}
		assert.GreaterOrEqual(t, len(entries), 1)
		entryUUID, err = uuid.Parse(entries[0])
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("test get a medialog entry", func(t *testing.T) {
		entry, err := mlc.GetEntryUUID(entryUUID)
		if err != nil {
			t.Error(err)
		}
		assert.Greater(t, entry.MediaID, uint(0))
	})
}
