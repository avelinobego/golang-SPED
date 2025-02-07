package schemes

import (
	"encoding/xml"
)

// ProcJudTrab ...
type ProcJudTrab struct {
	XMLName   xml.Name `xml:"procJudTrab"`
	NrProcJud string   `xml:"nrProcJud"`
	CodSusp   string   `xml:"codSusp"`
}
