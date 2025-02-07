package schemes

import (
	"encoding/xml"
)

// RendIsen0561 ...
type RendIsen0561 struct {
	XMLName               xml.Name `xml:"rendIsen0561"`
	VlrDiariasAttr        int      `xml:"vlrDiarias,attr,omitempty"`
	VlrAjudaCustoAttr     int      `xml:"vlrAjudaCusto,attr,omitempty"`
	VlrIndResContratoAttr int      `xml:"vlrIndResContrato,attr,omitempty"`
	VlrAbonoPecAttr       int      `xml:"vlrAbonoPec,attr,omitempty"`
	VlrAuxMoradiaAttr     int      `xml:"vlrAuxMoradia,attr,omitempty"`
}
