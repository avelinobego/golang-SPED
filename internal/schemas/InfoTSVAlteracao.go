package schemas

import (
	"encoding/xml"
)

// InfoTSVAlteracao ...
type InfoTSVAlteracao struct {
	XMLName            xml.Name            `xml:"infoTSVAlteracao"`
	DtAlteracao        string              `xml:"dtAlteracao"`
	NatAtividade       string              `xml:"natAtividade"`
	InfoComplementares *InfoComplementares `xml:"infoComplementares"`
}
