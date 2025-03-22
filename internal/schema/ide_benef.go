package schema

import (
	"encoding/xml"
)

// IdeBenef ...
type IdeBenef struct {
	XMLName       xml.Name         `xml:"ideBenef"`
	CpfBenef      string           `xml:"cpfBenef"`
	InfoPgto      []*InfoPgto      `xml:"infoPgto"`
	InfoIRComplem []*InfoIRComplem `xml:"infoIRComplem"`
}
