package schema

import (
	"encoding/xml"
)

// VerbasResc ...
type VerbasResc struct {
	XMLName     xml.Name `xml:"verbasResc"`
	DmDev       []*DmDev `xml:"dmDev"`
	ProcJudTrab []string `xml:"procJudTrab"`
	InfoMV      string   `xml:"infoMV"`
}
