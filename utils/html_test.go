package utils

import (
	"strings"
	"testing"

	"github.com/Woynert/course-mixer/samples"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestFindTagByName(t *testing.T) {
	for sampleName, sampleContents := range samples.Samples {
		t.Run(sampleName, func(tt *testing.T) {
			r := strings.NewReader(sampleContents)
			root, pErr := html.Parse(r)
			assert.Nil(tt, pErr)
			tBody := FindTagByName(root, "tbody")
			assert.NotNil(tt, tBody)
		})
	}
}
