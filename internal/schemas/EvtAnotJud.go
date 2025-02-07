package schemes

import (
	"encoding/xml"
)

// EvtAnotJud ...
type EvtAnotJud struct {
	XMLName       xml.Name      `xml:"evtAnotJud"`
	IdAttr        string        `xml:"Id,attr"`
	IdeEvento     string        `xml:"ideEvento"`
	IdeEmpregador string        `xml:"ideEmpregador"`
	InfoProcesso  *InfoProcesso `xml:"infoProcesso"`
	InfoAnotJud   *InfoAnotJud  `xml:"infoAnotJud"`
}
