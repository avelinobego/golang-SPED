package schemes

import (
	"encoding/xml"
)

// EvtReabreEvPer ...
type EvtReabreEvPer struct {
	XMLName       xml.Name `xml:"evtReabreEvPer"`
	IdAttr        string   `xml:"Id,attr"`
	IdeEvento     string   `xml:"ideEvento"`
	IdeEmpregador string   `xml:"ideEmpregador"`
}
