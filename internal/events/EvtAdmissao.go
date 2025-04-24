package events

import "encoding/xml"

type EvtAdmissao struct {
	XMLName xml.Name `xml:"xmlName"`
	Nome    string   `xml:"nome"`
	Idade   int32    `xml:"idade"`
}
