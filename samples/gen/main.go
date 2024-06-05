package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Woynert/course-mixer/parser"
	"github.com/Woynert/course-mixer/scraper"
	"github.com/Woynert/course-mixer/utils"
	"github.com/dave/jennifer/jen"
)

func downloadFile(sampleNumber int, form *scraper.Form, facultyName, faculty string) {
	log.Printf("[*] downloading %s\n", facultyName)
	res, err := scraper.DownloadFaculty(form.Action, faculty)
	utils.Fatal(err)
	defer res.Body.Close()
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, res.Body)
	utils.Fatal(err)
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
	utils.Fatal(err)
	defer sample.Close()
	io.Copy(sample, buffer2)
	log.Printf("[+] sample file created: %s\n", filename)
}

func downloadFiles() {
	form, err := scraper.GetForm(scraper.OfficialURL)
	utils.Fatal(err)
	sampleNumber := 1
	for facultyName, faculty := range form.Faculties {
		downloadFile(sampleNumber, form, facultyName, faculty)
		sampleNumber++
	}
}

func createSampleGo() {
	entries, err := os.ReadDir(".")
	utils.Fatal(err)
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
	utils.Fatal(err)
	defer samples.Close()
	_, err = samples.WriteString(fmt.Sprintf("%#v", file))
	utils.Fatal(err)
}

func main() {
	downloadFiles()
	createSampleGo()
}
