package utils

import (
	"strings"
	"testing"

	"github.com/PedroChaparro/PI202202-alako-data/samples"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestFindTagByName(t *testing.T) {
	t.Run("sample_1.html", func(tt *testing.T) {
		r := strings.NewReader(samples.Sample_1)
		root, pErr := html.Parse(r)
		assert.Nil(tt, pErr)
		tBody := FindTagByName(root, "tbody")
		assert.NotNil(tt, tBody)
	})
	t.Run("sample_2.html", func(tt *testing.T) {
		r := strings.NewReader(samples.Sample_2)
		root, pErr := html.Parse(r)
		assert.Nil(tt, pErr)
		tBody := FindTagByName(root, "tbody")
		assert.NotNil(tt, tBody)
	})
}
