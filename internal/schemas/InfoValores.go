package schemes

import (
	"encoding/xml"
)

// InfoValores ...
type InfoValores struct {
	XMLName      xml.Name   `xml:"infoValores"`
	IndApuracao  string     `xml:"indApuracao"`
	VlrNRetido   string     `xml:"vlrNRetido"`
	VlrDepJud    string     `xml:"vlrDepJud"`
	VlrCmpAnoCal string     `xml:"vlrCmpAnoCal"`
	VlrCmpAnoAnt string     `xml:"vlrCmpAnoAnt"`
	VlrRendSusp  string     `xml:"vlrRendSusp"`
	DedSusp      []*DedSusp `xml:"dedSusp"`
}
