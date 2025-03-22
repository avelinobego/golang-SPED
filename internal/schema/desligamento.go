package schema

import (
	"encoding/xml"
)

// Desligamento ...
type Desligamento struct {
	XMLName      xml.Name `xml:"desligamento"`
	MtvDeslig    string   `xml:"mtvDeslig"`
	DtDeslig     string   `xml:"dtDeslig"`
	DtProjFimAPI string   `xml:"dtProjFimAPI"`
}
