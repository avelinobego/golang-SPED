package schemas

import (
	"encoding/xml"
)

// AgNoc ...
type AgNoc struct {
	XMLName    xml.Name `xml:"agNoc"`
	CodAgNoc   string   `xml:"codAgNoc"`
	DscAgNoc   string   `xml:"dscAgNoc"`
	TpAval     int8     `xml:"tpAval"`
	IntConc    float64  `xml:"intConc"`
	LimTol     float64  `xml:"limTol"`
	UnMed      int8     `xml:"unMed"`
	TecMedicao string   `xml:"tecMedicao"`
	NrProcJud  string   `xml:"nrProcJud"`
	EpcEpi     *EpcEpi  `xml:"epcEpi"`
}
