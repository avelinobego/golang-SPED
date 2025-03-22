package schema

import (
	"encoding/xml"
)

// InfoDeslig ...
type InfoDeslig struct {
	XMLName      xml.Name `xml:"infoDeslig"`
	DtDeslig     string   `xml:"dtDeslig"`
	MtvDeslig    string   `xml:"mtvDeslig"`
	DtProjFimAPI string   `xml:"dtProjFimAPI"`
	PensAlim     string   `xml:"pensAlim"`
	PercAliment  string   `xml:"percAliment"`
	VrAlim       string   `xml:"vrAlim"`
}
