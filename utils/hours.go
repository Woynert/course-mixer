package utils

import (
	"fmt"
	"regexp"
	"time"
)

var (
	timeFormat      = "03:04 PM"
	timeRegexp      = regexp.MustCompile(`\d{2}:\d{2} (AM|PM)`)
	classroomRegexp = regexp.MustCompile(`[A-Za-z]\d+`)
)

type Hour struct {
	Day       time.Weekday `json:"day"`
	From      time.Time    `json:"from"`
	To        time.Time    `json:"to"`
	Classroom string       `json:"classroom"`
}

func ParseHours(day time.Weekday, hour string) (*Hour, error) {
	times := timeRegexp.FindAllString(hour, -1)
	if len(times) != 2 {
		return nil, fmt.Errorf("expecting format such as: 12:00 PM-01:40 PM C203 but received: %s", hour)
	}
	// The previous regexp makes sure this always is OK
	var from, to time.Time
	from, _ = time.Parse(timeFormat, times[0])
	to, _ = time.Parse(timeFormat, times[1])
	classroom := classroomRegexp.FindString(hour)
	if classroom == "" {
		classroom = "unknown"
	}
	class := &Hour{
		Day:       day,
		From:      from,
		To:        to,
		Classroom: classroom,
	}
	return class, nil
}
