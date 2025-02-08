package schemas

import (
	"encoding/xml"
)

// InfoExclusao ...
type InfoExclusao struct {
	XMLName     xml.Name     `xml:"infoExclusao"`
	TpEvento    string       `xml:"tpEvento"`
	NrRecEvt    string       `xml:"nrRecEvt"`
	IdeProcTrab *IdeProcTrab `xml:"ideProcTrab"`
}
