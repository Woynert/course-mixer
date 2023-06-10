package main

import (
	"encoding/json" // Write json file
	"fmt"           // Parse requests
	"os"

	//"bufio"

	// Regular expressions

	"github.com/PedroChaparro/PI202202-alako-data/scraper"
	"github.com/PedroChaparro/PI202202-alako-data/utils"
)

type Query struct {
	Name string `json:"name"`
	File string `json:"file"`
	Cols uint8  `json:"cols"`
}

type Hour struct {
	Day   string `json:"day"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

type Course struct {
	Ctg   string `json:"ctg"`
	Level uint8  `json:"level"`
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

func createNewHour(day string, hourStr string) Hour {
	fmt.Println(day, hourStr)
	newH := Hour{
		Day:   "lun",
		Start: 10,
		End:   11,
	}
	panic(1)
	return newH
}

func downloadData() []Course {
	courses, err := scraper.Crawl()
	utils.Fatal(err)
	result := make([]Course, 0, len(courses))
	for _, course := range courses {
		hours := make([]Hour, 0, len(course.Hours))
		for _, hour := range course.Hours {
			h := Hour{
				Day:   hour.Day.String(),
				Start: hour.From.Hour(),
				End:   hour.To.Hour(),
			}
			hours = append(hours, h)
		}
		c := Course{
			Ctg:   course.Faculty,
			Level: course.Level,
			Title: course.Subject,
			Nrc:   course.NRC,
			Hours: hours,
		}
		result = append(result, c)
	}
	return result
}

func main() {
	output := os.Stdout
	if len(os.Args) > 1 {
		var err error
		output, err = os.OpenFile(os.Args[1], os.O_WRONLY, 0777)
		utils.Fatal(err)
	}
	courses := downloadData()
	encoder := json.NewEncoder(output)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(courses)
	utils.Fatal(err)

}
