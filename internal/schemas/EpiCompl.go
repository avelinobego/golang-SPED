package schemas

import (
	"encoding/xml"
)

// EpiCompl ...
type EpiCompl struct {
	XMLName       xml.Name `xml:"epiCompl"`
	MedProtecao   string   `xml:"medProtecao"`
	CondFuncto    string   `xml:"condFuncto"`
	UsoInint      string   `xml:"usoInint"`
	PrzValid      string   `xml:"przValid"`
	PeriodicTroca string   `xml:"periodicTroca"`
	Higienizacao  string   `xml:"higienizacao"`
}
