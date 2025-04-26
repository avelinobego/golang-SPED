package events

import "encoding/xml"

type EvtAdmissao struct {
	XMLName xml.Name `xml:"evtAdmissao"`
	Nome    string   `xml:"nome"`
	Idade   int32    `xml:"idade"`
}
