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

//go:embed sample_15.html
var Sample_15 string

//go:embed sample_16.html
var Sample_16 string

//go:embed sample_17.html
var Sample_17 string

//go:embed sample_18.html
var Sample_18 string

//go:embed sample_2.html
var Sample_2 string

//go:embed sample_3.html
var Sample_3 string

//go:embed sample_4.html
var Sample_4 string

//go:embed sample_5.html
var Sample_5 string

//go:embed sample_8.html
var Sample_8 string

//go:embed sample_9.html
var Sample_9 string

var Samples = map[string]string{
	"Sample_1":  Sample_1,
	"Sample_10": Sample_10,
	"Sample_11": Sample_11,
	"Sample_12": Sample_12,
	"Sample_13": Sample_13,
	"Sample_14": Sample_14,
	"Sample_15": Sample_15,
	"Sample_16": Sample_16,
	"Sample_17": Sample_17,
	"Sample_18": Sample_18,
	"Sample_2":  Sample_2,
	"Sample_3":  Sample_3,
	"Sample_4":  Sample_4,
	"Sample_5":  Sample_5,
	"Sample_8":  Sample_8,
	"Sample_9":  Sample_9,
}
