package schema

import (
	"encoding/xml"
)

// InfoTrabCedido ...
type InfoTrabCedido struct {
	XMLName   xml.Name `xml:"infoTrabCedido"`
	CategOrig string   `xml:"categOrig"`
	CnpjCednt string   `xml:"cnpjCednt"`
	MatricCed string   `xml:"matricCed"`
	DtAdmCed  string   `xml:"dtAdmCed"`
	TpRegTrab string   `xml:"tpRegTrab"`
	TpRegPrev string   `xml:"tpRegPrev"`
}
