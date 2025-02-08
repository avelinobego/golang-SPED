package schemas

import (
	"encoding/xml"
)

// IdeResp ...
type IdeResp struct {
	XMLName      xml.Name `xml:"ideResp"`
	TpInsc       string   `xml:"tpInsc"`
	NrInsc       string   `xml:"nrInsc"`
	DtAdmRespDir string   `xml:"dtAdmRespDir"`
	MatRespDir   string   `xml:"matRespDir"`
}
