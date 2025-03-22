package schema

import (
	"encoding/xml"
)

// InfoIRComplem ...
type InfoIRComplem struct {
	XMLName      xml.Name        `xml:"infoIRComplem"`
	DtLaudo      string          `xml:"dtLaudo"`
	PerAnt       *PerAnt         `xml:"perAnt"`
	InfoDep      []*InfoDep      `xml:"infoDep"`
	InfoIRCR     []*InfoIRCR     `xml:"infoIRCR"`
	PlanSaude    []*PlanSaude    `xml:"planSaude"`
	InfoReembMed []*InfoReembMed `xml:"infoReembMed"`
}
