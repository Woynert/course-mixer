package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseHours(t *testing.T) {
	t.Run("Must: 12:00 PM-01:40 PM C203", func(tt *testing.T) {
		expect := Hour{
			Day:       time.Monday,
			From:      time.Date(0, time.January, 1, 12, 0, 0, 0, time.UTC),
			To:        time.Date(0, time.January, 1, 13, 40, 0, 0, time.UTC),
			Classroom: "C203",
		}
		class, err := ParseHours(time.Monday, "12:00 PM-01:40 PM C203")
		assert.Nil(tt, err)
		assert.Equal(tt, expect, *class)
		assert.Equal(tt, 12, class.From.Hour())
		assert.Equal(tt, 13, class.To.Hour())
	})
	t.Run("Invalid Format", func(tt *testing.T) {
		_, err := ParseHours(time.Monday, "12:000 PM-01:40 PM C203")
		assert.NotNil(tt, err)
	})
	t.Run("No Classroom", func(tt *testing.T) {
		expect := Hour{
			Day:       time.Monday,
			From:      time.Date(0, time.January, 1, 12, 0, 0, 0, time.UTC),
			To:        time.Date(0, time.January, 1, 13, 40, 0, 0, time.UTC),
			Classroom: "unknown",
		}
		class, err := ParseHours(time.Monday, "12:00 PM-01:40 PM")
		assert.Nil(tt, err)
		assert.Equal(tt, expect, *class)
		assert.Equal(tt, 12, class.From.Hour())
		assert.Equal(tt, 13, class.To.Hour())
	})
}
