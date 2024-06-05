package scraper

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"log"

	"github.com/Woynert/course-mixer/parser"
	"github.com/Woynert/course-mixer/utils"
	"golang.org/x/net/html"
)

const OfficialURL = "https://horariosupb.bucaramanga.upb.edu.co/horariosclases"

type Form struct {
	Action    string            `json:"action"`
	Faculties map[string]string `json:"faculties"`
}

func GetForm(classesUrl string) (*Form, error) {
	res, err := http.Get(classesUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %w", classesUrl, err)
	}
	defer res.Body.Close()
	root, err := html.Parse(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not parse %s html: %w", classesUrl, err)
	}
	f := utils.FindTagByName(root, "form")
	if f == nil {
		return nil, fmt.Errorf("form not found in %s", classesUrl)
	}
	var action string
	for _, attr := range f.Attr {
		if attr.Key == "action" {
			action = attr.Val
		}
	}
	if action == "" {
		return nil, fmt.Errorf("invalid form: no action found")
	}
	s := utils.FindTagByName(root, "select")
	if s == nil {
		return nil, fmt.Errorf("select not found in %s", classesUrl)
	}
	form := &Form{
		Action:    action,
		Faculties: map[string]string{},
	}
	for option := s.FirstChild; option != nil; option = option.NextSibling {
		for _, attr := range option.Attr {
			if attr.Key == "value" && attr.Val != "" {
				form.Faculties[option.FirstChild.Data] = attr.Val
				break
			}
		}
	}
	return form, nil
}

var (
	tabRegexp    = regexp.MustCompile(`\$\("#i-clases"\)\.click\(function\(\)\{\s+\S+`)
	actionRegexp = regexp.MustCompile(`"\?var=\S+"`)
)

func DownloadFaculty(action, facultyValue string) (*http.Response, error) {
	data := url.Values{
		"n_idfacultad": []string{facultyValue},
	}
	fullURL := fmt.Sprintf("%s?%s", action, data.Encode())
	return http.Get(fullURL)
}

func ParseFaculty(action, facultyValue string) ([]*parser.Course, error) {
	res, err := DownloadFaculty(action, facultyValue)
	if err != nil {
		return nil, fmt.Errorf("error while requesting courses for %s: %w", facultyValue, err)
	}
	defer res.Body.Close()
	return parser.Parse(res.Body)
}

func Crawl() ([]*parser.Course, error) {
	// First Get request to obtain a valid Form token, the value of the select
	form, err := GetForm(OfficialURL)
	if err != nil {
		return nil, fmt.Errorf("could not get action: %w", err)
	}
	var courses []*parser.Course
	for facultyName, facultyValue := range form.Faculties {
		result, err := ParseFaculty(form.Action, facultyValue)
		if err != nil {
			log.Printf("Error %s: %s", facultyName, err.Error())
			continue
		}
		log.Printf("Done  %s", facultyName)
		courses = append(courses, result...)
	}
	return courses, nil
}
