package schemes

import (
	"encoding/xml"
)

// ProcAdmJudFap ...
type ProcAdmJudFap struct {
	XMLName xml.Name `xml:"procAdmJudFap"`
	TpProc  string   `xml:"tpProc"`
	NrProc  string   `xml:"nrProc"`
	CodSusp string   `xml:"codSusp"`
}
