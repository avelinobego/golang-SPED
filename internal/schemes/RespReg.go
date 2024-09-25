package schemes

import (
	"encoding/xml"
)

// RespReg ...
type RespReg struct {
	XMLName xml.Name `xml:"respReg"`
	CpfResp string   `xml:"cpfResp"`
	IdeOC   int8     `xml:"ideOC"`
	DscOC   string   `xml:"dscOC"`
	NrOC    string   `xml:"nrOC"`
	UfOC    string   `xml:"ufOC"`
}
