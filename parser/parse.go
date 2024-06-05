package parser

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"time"

	"github.com/Woynert/course-mixer/utils"
	"golang.org/x/net/html"
)

var (
	ErrInvalidHTML = fmt.Errorf("invalid HTML structure")
)

type Course struct {
	Faculty      string        `json:"faculty"`
	Subject      string        `json:"subject"`
	Matter       string        `json:"matter"`
	CourseNumber string        `json:"courseNumber"`
	NRC          string        `json:"nrc"`
	Hours        []*utils.Hour `json:"hours"`
	Credits      uint8         `json:"credits"`
	Level        uint8         `json:"level"`
}

func withoutLevel(children []string) (c *Course, err error) {
	var credits uint64
	credits, err = strconv.ParseUint(children[5], 10, 8)
	if err == nil {
		c = &Course{
			Faculty:      children[0],
			Level:        0,
			Subject:      children[1],
			Matter:       children[2],
			CourseNumber: children[3],
			NRC:          children[4],
			Credits:      uint8(credits),
		}
		var hour *utils.Hour
		for dayIndex, hourStr := range children[6:] {
			if hourStr != "\u00a0" {
				hour, err = utils.ParseHours(time.Weekday(dayIndex+1), hourStr)
				if err == nil {
					c.Hours = append(c.Hours, hour)
				}
			}
		}
	}
	return c, nil
}

func withLevel(children []string) (c *Course, err error) {
	var (
		credits uint64
		level   uint64
	)
	level, err = strconv.ParseUint(children[1], 10, 8)
	if err == nil {
		credits, err = strconv.ParseUint(children[6], 10, 8)
		if err == nil {
			c = &Course{
				Faculty:      children[0],
				Subject:      children[2],
				Matter:       children[3],
				CourseNumber: children[4],
				NRC:          children[5],
				Level:        uint8(level),
				Credits:      uint8(credits),
			}
			var hour *utils.Hour
			for dayIndex, hourStr := range children[7:] {
				if hourStr != "" {
					hour, err = utils.ParseHours(time.Weekday(dayIndex+1), hourStr)
					if err == nil {
						c.Hours = append(c.Hours, hour)
					}
				}
			}
		}
	}
	return c, nil
}

var number = regexp.MustCompile(`(?m)^\d+$`)

func parseCourse(tr *html.Node) (*Course, error) {
	children := make([]string, 0, 13)
	for td := tr.FirstChild; td != nil; td = td.NextSibling {
		if td.Data != "td" || td.FirstChild == nil {
			continue
		}
		tdChild := td.FirstChild.NextSibling
		if tdChild == nil || tdChild.FirstChild == nil {
			children = append(children, "")
			continue
		}
		children = append(children, tdChild.FirstChild.Data)
	}
	switch {
	case len(children) <= 6:
		return nil, nil
	case number.MatchString(children[1]):
		return withLevel(children)
	default:
		return withoutLevel(children)
	}
}

func Parse(r io.Reader) ([]*Course, error) {
	root, pErr := html.Parse(r)
	if pErr != nil {
		return nil, fmt.Errorf("error while parsing: %w", pErr)
	}
	tBody := utils.FindTagByName(root, "tbody")
	if tBody == nil {
		return nil, fmt.Errorf("no tbody found: %w", ErrInvalidHTML)
	}
	courses := make([]*Course, 0, 40)
	for tr := tBody.FirstChild; tr != nil; tr = tr.NextSibling {
		if tr.Data != "tr" {
			continue
		}
		course, pErr := parseCourse(tr)
		if pErr != nil {
			return nil, pErr
		}
		if course != nil {
			courses = append(courses, course)
		}
	}
	return courses, nil
}
