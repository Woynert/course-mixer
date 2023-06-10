package parser

import (
	"strings"
	"testing"

	"github.com/Woynert/course-mixer/samples"
	"github.com/Woynert/course-mixer/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func Test_parseCourse(t *testing.T) {
	for sampleName, sampleContents := range samples.Samples {
		t.Run(sampleName, func(tt *testing.T) {
			r := strings.NewReader(sampleContents)
			root, pErr := html.Parse(r)
			assert.Nil(tt, pErr)
			tBody := utils.FindTagByName(root, "tbody")
			assert.NotNil(tt, tBody)
			var courses []*Course
			for tr := tBody.FirstChild; tr != nil; tr = tr.NextSibling {
				if tr.Data != "tr" {
					continue
				}
				course, pErr := parseCourse(tr)
				assert.Nil(tt, pErr)
				if course != nil {
					courses = append(courses, course)
				}
			}
			assert.NotEmpty(tt, courses)
		})
	}
}

func TestParse(t *testing.T) {
	for sampleName, sampleContents := range samples.Samples {
		t.Run(sampleName, func(tt *testing.T) {
			r := strings.NewReader(sampleContents)
			courses, pErr := Parse(r)
			assert.Nil(tt, pErr)
			assert.Greater(tt, len(courses), 0)
		})
	}
}
