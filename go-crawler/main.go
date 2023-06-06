package main

import (
	"fmt"
	"io/ioutil" // Parse requests
	"encoding/json" // Write json file
	"os"
	//"bufio"
	"time"
	"regexp" // Regular expressions
	"strings"
	"bytes"
	"strconv"

	
	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)

type Query struct {
	Name string `json:"name"`	
	File string `json:"file"`
	Cols int `json:"cols"`
}

type Hour struct {
	Day   string `json:"day"`
	Start int `json:"start"`
	End   int `json:"end"`
}

type Course struct {
	Ctg string `json:"ctg"`
	Level string `json:"level"`
	Title string `json:"title"`
	Nrc   string `json:"nrc"`
	Hours []Hour `json:"hour"`
}

// NOTE: Define your queries here
// Name - Optional
// File - The html file name inside the datain folder
// Cols - The amount of columns found

var queries = []Query{
	{
		Name: "Sistemas",
		File: "sistemas.html",
		Cols: 13,
	},
	{
		Name: "Electivas",
		File: "electivas.html",
		Cols: 12,
	},
}

var courses []Course = []Course{}
var pathWD string

// TODO handle this case 12:00 PM-01:40 PM C203
func formatHours(clnStr string) (int, int) {
	fmt.Println("UUUUUUUUUUUUUUUU", clnStr, "UU")
	words := strings.Split(clnStr, "-")

	// first digit
	firDig, _ := strconv.Atoi(strings.Split(words[0], ":")[0])
	firAmpm := strings.Split(words[0], " ")[1]

	// second
	secDig, _ := strconv.Atoi(strings.Split(words[1], ":")[0])
	secAmpm := strings.Split(words[1], " ")[1]

	if (firAmpm == "PM" && firDig < 12){
		firDig += 12
	}

	if (secAmpm == "PM"){
		secDig += 12
	}

	secDig++

	fmt.Println(firDig, firAmpm)
	fmt.Println(secDig, secAmpm)
	//return strconv.Itoa(firDig), strconv.Itoa(secDig)
	return firDig, secDig
}

func createNewHour(day string, hourStr string) Hour{

	// create new hours
	newhstart, newhend := formatHours(hourStr)
	newh := Hour{
		Day  : day,
		Start: newhstart,
		End  : newhend,
	}

	fmt.Println(newh)
	return newh
}

func getData(file string, cols int, ctg string) (err error) {

	fmt.Printf("Current file: %s\n", file)

	// ### ### ###
	// Get plain html
	// Open the file.
	f, _ := os.Open(file)
	defer f.Close()

	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	bodyStr := string(b)

	// regex

	regexTd, err := regexp.Compile("<[^>]+>([^<]+)")
	if err != nil {
		fmt.Println("Invalid regexp")
		return // problem with the regular expression.
	}

	// css select

	selectorStr := fmt.Sprintf("td:first-child:nth-last-child(%d) ~ td", cols)
	fmt.Println(selectorStr)

	sel, err := css.Parse(selectorStr)
	if err != nil {
		panic(err)
	}
	node, err := html.Parse(strings.NewReader(bodyStr))
	if err != nil {
		panic(err)
	}

	// filter lines

	counter := 0
	skip := true

	offset := 0
	if (cols > 12) {
		offset = 1
	}

	// course properties
	course := Course{
		Level: "",
		Title: "",
		Nrc  : "",
		Hours: []Hour{},
	}

	for _, ele := range sel.Select(node) {


		var strBuff bytes.Buffer
		html.Render(&strBuff, ele)

		fmt.Println(counter -offset, "===", skip)

		if (counter == 0){
			//if (!skip){
				//fmt.Println(course)
			//}

			skip = false
			course = Course{
				Level: "",
				Title: "",
				Nrc  : "",
				Hours: []Hour{},
			}
		}

		if (!skip){

			var clnStr string = ""
			rawStr := strBuff.String()
			submatch := regexTd.FindStringSubmatch(rawStr)
			if (len(submatch) > 1){
				clnStr = submatch[1]
			}

			fmt.Println(clnStr)

			
			switch (counter-offset) {

				// title
				case 0:

					if (rawStr == "<td>ASIGNATURA</td>" || rawStr == "<td> </td>"){
						skip = true
					} else {
						course.Title = clnStr
					}

				case -1:
					
					// it has levels
					if (cols > 12) {
						course.Level = clnStr
					}

				case 3:
					course.Nrc = clnStr

				// lunes
				case 5:
					if (strings.TrimSpace(clnStr) != ""){
						course.Hours = append(course.Hours, createNewHour("lun", clnStr))
					}
				case 6:
					if (strings.TrimSpace(clnStr) != ""){
						course.Hours = append(course.Hours, createNewHour("mar", clnStr))
					}
				case 7:
					if (strings.TrimSpace(clnStr) != ""){
						course.Hours = append(course.Hours, createNewHour("mie", clnStr))
					}
				case 8:
					if (strings.TrimSpace(clnStr) != ""){
						course.Hours = append(course.Hours, createNewHour("jue", clnStr))
					}
				case 9:
					if (strings.TrimSpace(clnStr) != ""){
						course.Hours = append(course.Hours, createNewHour("vie", clnStr))
					}
				case 10:
					if (strings.TrimSpace(clnStr) != ""){
						course.Hours = append(course.Hours, createNewHour("sab", clnStr))
					}
					if (!skip){
						course.Ctg = ctg
						fmt.Println(course)
						courses = append(courses, course)
					}

			}

			//fmt.Println(str)

		}

		counter++
		counter = counter % (cols -1)
	}


	return nil
}

func main(){

	pathWD, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	// For each query
	for _, query := range(queries) {

		currLength := len(courses)
		start := time.Now()
		filePath := pathWD + "/datain/" + query.File

		fmt.Printf("🏃 Starting with query: %s\n", query.File)
		getData(filePath, query.Cols, query.Name)
		fmt.Println(courses)

		fmt.Printf("🏁 %d items were saved in %s\n\n", len(courses) - currLength, time.Since(start))
	}

	// Create json file
	jsonString, _ := json.MarshalIndent(courses, "", "	")
	ioutil.WriteFile("dataout/data.json", jsonString, os.ModePerm)

}
