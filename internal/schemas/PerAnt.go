package schemas

import (
	"encoding/xml"
)

// PerAnt ...
type PerAnt struct {
	XMLName       xml.Name `xml:"perAnt"`
	PerRefAjuste  string   `xml:"perRefAjuste"`
	NrRec1210Orig string   `xml:"nrRec1210Orig"`
}
