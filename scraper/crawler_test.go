package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetForm(t *testing.T) {
	form, err := GetForm(OfficialURL)
	assert.Nil(t, err)
	assert.NotEmpty(t, form.Action)
	assert.NotEmpty(t, form.Faculties)
}

func TestCrawl(t *testing.T) {
	courses, err := Crawl()
	assert.Nil(t, err)
	assert.NotEmpty(t, courses)
}
