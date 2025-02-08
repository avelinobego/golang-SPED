package schemas

import (
	"encoding/xml"
)

// CalcTrib ...
type CalcTrib struct {
	XMLName          xml.Name `xml:"calcTrib"`
	PerRefAttr       int      `xml:"perRef,attr"`
	VrBcCpMensalAttr int      `xml:"vrBcCpMensal,attr"`
	VrBcCp13Attr     int      `xml:"vrBcCp13,attr"`
	IdeSeqProc       int      `xml:"ideSeqProc"`
}
