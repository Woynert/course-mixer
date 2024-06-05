package samples

import _ "embed"

//go:embed sample_1.html
var Sample_1 string

//go:embed sample_10.html
var Sample_10 string

//go:embed sample_11.html
var Sample_11 string

//go:embed sample_12.html
var Sample_12 string

//go:embed sample_13.html
var Sample_13 string

//go:embed sample_14.html
var Sample_14 string

//go:embed sample_18.html
var Sample_18 string

//go:embed sample_3.html
var Sample_3 string

//go:embed sample_6.html
var Sample_6 string

//go:embed sample_7.html
var Sample_7 string

var Samples = map[string]string{
	"Sample_1":  Sample_1,
	"Sample_10": Sample_10,
	"Sample_11": Sample_11,
	"Sample_12": Sample_12,
	"Sample_13": Sample_13,
	"Sample_14": Sample_14,
	"Sample_18": Sample_18,
	"Sample_3":  Sample_3,
	"Sample_6":  Sample_6,
	"Sample_7":  Sample_7,
}
