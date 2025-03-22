package schema

import (
	"encoding/xml"
)

// InstPenMorte ...
type InstPenMorte struct {
	XMLName xml.Name `xml:"instPenMorte"`
	CpfInst string   `xml:"cpfInst"`
	DtInst  string   `xml:"dtInst"`
}
