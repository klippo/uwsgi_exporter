package exporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileStatsReader_Read(t *testing.T) {
	a := assert.New(t)

	uri := "file://" + sampleUwsgiStatsFileName

	reader, err := NewStatsReader(uri, someTimeout)
	a.NoError(err)

	uwsgiStats, err := reader.Read()
	a.NoError(err)

	a.Equal(uwsgiStats.Version, "2.0.12")
}
