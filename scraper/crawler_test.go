package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClassesUrl(t *testing.T) {
	classesUrl, err := GetClassesUrl(OfficialURL)
	assert.Nil(t, err)
	assert.NotEmpty(t, classesUrl)
}

func TestGetForm(t *testing.T) {
	classesUrl, err := GetClassesUrl(OfficialURL)
	assert.Nil(t, err)
	assert.NotEmpty(t, classesUrl)
	form, err := GetForm(OfficialURL, classesUrl)
	assert.Nil(t, err)
	assert.NotEmpty(t, form.Action)
	assert.NotEmpty(t, form.Faculties)
}

func TestCrawl(t *testing.T) {
	courses, err := Crawl()
	assert.Nil(t, err)
	assert.NotEmpty(t, courses)
}
