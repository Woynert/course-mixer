package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/PedroChaparro/PI202202-alako-data/parser"
	"github.com/PedroChaparro/PI202202-alako-data/scraper"
	"github.com/dave/jennifer/jen"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func downloadFile(sampleNumber int, form *scraper.Form, facultyName, faculty string) {
	log.Printf("[*] downloading %s\n", facultyName)
	res, err := scraper.DownloadFaculty(form.Action, faculty)
	fatal(err)
	defer res.Body.Close()
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, res.Body)
	fatal(err)
	buffer2 := bytes.NewBuffer(buffer.Bytes())
	courses, err := parser.Parse(&buffer)
	if err != nil {
		log.Printf("[!] Failed parse for %s: %s\n", facultyName, err.Error())
		return
	}
	if courses == nil {
		log.Printf("[!] Failed parse for %s: no courses found\n", facultyName)
		return
	}
	filename := fmt.Sprintf("sample_%d.html", sampleNumber)
	log.Printf("[*] writing file to: %s\n", filename)
	sample, err := os.Create(filename)
	fatal(err)
	defer sample.Close()
	io.Copy(sample, buffer2)
	log.Printf("[+] sample file created: %s\n", filename)
}

func downloadFiles() {
	classesUrl, err := scraper.GetClassesUrl(scraper.OfficialURL)
	fatal(err)
	form, err := scraper.GetForm(scraper.OfficialURL, classesUrl)
	fatal(err)
	sampleNumber := 1
	for facultyName, faculty := range form.Faculties {
		downloadFile(sampleNumber, form, facultyName, faculty)
		sampleNumber++
	}
}

func createSampleGo() {
	entries, err := os.ReadDir(".")
	fatal(err)
	file := jen.NewFile("samples")
	file.Anon("embed")
	var variables jen.Dict = jen.Dict{}
	for _, entry := range entries {
		if entry.IsDir() || !strings.Contains(entry.Name(), ".html") {
			continue
		}
		splitName := strings.Split(entry.Name(), ".")
		variable := strings.Title(splitName[0])
		variables[jen.Lit(variable)] = jen.Id(variable)
		file.
			Comment(fmt.Sprintf("//go:embed %s", entry.Name())).
			Line().Var().Id(variable).String()
	}
	file.Line().Var().Id("Samples").Op("=").Map(jen.String()).String().Values(variables)
	samples, err := os.Create("samples.go")
	fatal(err)
	defer samples.Close()
	_, err = samples.WriteString(fmt.Sprintf("%#v", file))
	fatal(err)
}

func main() {
	downloadFiles()
	createSampleGo()
}
