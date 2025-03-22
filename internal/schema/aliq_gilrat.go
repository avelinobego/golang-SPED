package schema

import (
	"encoding/xml"
)

// AliqGilrat is Se informado, deve ser um número maior ou igual a 0,5000 e menor ou igual a 2,0000 e, no caso da alínea "b", deve ser diferente do valor definido pelo órgão governamental competente.
type AliqGilrat struct {
	XMLName       xml.Name       `xml:"aliqGilrat"`
	AliqRat       string         `xml:"aliqRat"`
	Fap           string         `xml:"fap"`
	ProcAdmJudRat *ProcAdmJudRat `xml:"procAdmJudRat"`
	ProcAdmJudFap *ProcAdmJudFap `xml:"procAdmJudFap"`
}
