package schema

import "encoding/xml"

type EvtInfoEmpregador struct {
	XMLName xml.Name `xml:"EvtInfoEmpregador"`
	Id      string   `xml:Id,attr`
}
