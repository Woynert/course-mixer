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

const OfficialURL = "https://horariosupb.bucaramanga.upb.edu.co/"

type Form struct {
	Action    string            `json:"action"`
	Faculties map[string]string `json:"faculties"`
}

func GetForm(target, classesUrl string) (*Form, error) {
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
		Action:    target + action,
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

func GetClassesUrl(target string) (finalUrl string, err error) {
	res, err := http.Get(target)
	if err != nil {
		return "", fmt.Errorf("failed to get %s: %w", target, err)
	}
	defer res.Body.Close()
	root, err := html.Parse(res.Body)
	if err != nil {
		return "", fmt.Errorf("could not parse %s html: %w", target, err)
	}
	body := utils.FindTagByName(root, "body")
	if body == nil {
		return "", fmt.Errorf("body not found in %s", target)
	}
	script := utils.FindTagByName(body, "script")
	if script == nil {
		return "", fmt.Errorf("script not found in %s", target)
	}
	iClass := tabRegexp.FindString(script.FirstChild.Data)
	actionPage := actionRegexp.FindString(iClass)
	actionPage = actionPage[1 : len(actionPage)-1]
	return target + actionPage, nil
}

func DownloadFaculty(action, faculty string) (*http.Response, error) {
	data := url.Values{
		"facultad": []string{faculty},
	}
	return http.PostForm(action, data)
}

func ParseFaculty(action, faculty string) ([]*parser.Course, error) {
	res, err := DownloadFaculty(action, faculty)
	if err != nil {
		return nil, fmt.Errorf("error while requesting courses for %s: %w", faculty, err)
	}
	defer res.Body.Close()
	return parser.Parse(res.Body)
}

func Crawl() ([]*parser.Course, error) {
	classesUrl, err := GetClassesUrl(OfficialURL)
	if err != nil {
		return nil, fmt.Errorf("could not obtain classes url: %w", err)
	}
	// First Get request to obtain a valid Form token, the value of the select
	form, err := GetForm(OfficialURL, classesUrl)
	if err != nil {
		return nil, fmt.Errorf("could not get action: %w", err)
	}
	var courses []*parser.Course
	for facultyName, faculty := range form.Faculties {
		result, err := ParseFaculty(form.Action, faculty)
		if err != nil {
			log.Printf("Error %s: %s", facultyName, err.Error())
			continue
			// return nil, fmt.Errorf("error while downloading faculty: %s: %w", faculty, err)
		}
		courses = append(courses, result...)
	}
	return courses, nil
}
