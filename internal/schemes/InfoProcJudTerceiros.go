package schemes

import (
	"encoding/xml"
)

// InfoProcJudTerceiros ...
type InfoProcJudTerceiros struct {
	XMLName         xml.Name           `xml:"infoProcJudTerceiros"`
	ProcJudTerceiro []*ProcJudTerceiro `xml:"procJudTerceiro"`
}
