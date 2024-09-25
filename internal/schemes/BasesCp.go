package schemes

import (
	"encoding/xml"
)

// BasesCp ...
type BasesCp struct {
	XMLName        xml.Name `xml:"basesCp"`
	VrBcCp00       string   `xml:"vrBcCp00"`
	VrBcCp15       string   `xml:"vrBcCp15"`
	VrBcCp20       string   `xml:"vrBcCp20"`
	VrBcCp25       string   `xml:"vrBcCp25"`
	VrSuspBcCp00   string   `xml:"vrSuspBcCp00"`
	VrSuspBcCp15   string   `xml:"vrSuspBcCp15"`
	VrSuspBcCp20   string   `xml:"vrSuspBcCp20"`
	VrSuspBcCp25   string   `xml:"vrSuspBcCp25"`
	VrBcCp00VA     string   `xml:"vrBcCp00VA"`
	VrBcCp15VA     string   `xml:"vrBcCp15VA"`
	VrBcCp20VA     string   `xml:"vrBcCp20VA"`
	VrBcCp25VA     string   `xml:"vrBcCp25VA"`
	VrSuspBcCp00VA string   `xml:"vrSuspBcCp00VA"`
	VrSuspBcCp15VA string   `xml:"vrSuspBcCp15VA"`
	VrSuspBcCp20VA string   `xml:"vrSuspBcCp20VA"`
	VrSuspBcCp25VA string   `xml:"vrSuspBcCp25VA"`
	VrDescSest     string   `xml:"vrDescSest"`
	VrCalcSest     string   `xml:"vrCalcSest"`
	VrDescSenat    string   `xml:"vrDescSenat"`
	VrCalcSenat    string   `xml:"vrCalcSenat"`
	VrSalFam       string   `xml:"vrSalFam"`
	VrSalMat       string   `xml:"vrSalMat"`
}
