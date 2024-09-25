package schemes

import (
	"encoding/xml"
)

// InfoDirigenteSindical ...
type InfoDirigenteSindical struct {
	XMLName    xml.Name `xml:"infoDirigenteSindical"`
	CategOrig  string   `xml:"categOrig"`
	TpInsc     string   `xml:"tpInsc"`
	NrInsc     string   `xml:"nrInsc"`
	DtAdmOrig  string   `xml:"dtAdmOrig"`
	MatricOrig string   `xml:"matricOrig"`
	TpRegTrab  string   `xml:"tpRegTrab"`
	TpRegPrev  string   `xml:"tpRegPrev"`
}
