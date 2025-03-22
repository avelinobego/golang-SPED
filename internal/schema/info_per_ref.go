package schema

import (
	"encoding/xml"
)

// InfoPerRef ...
type InfoPerRef struct {
	XMLName       xml.Name         `xml:"infoPerRef"`
	PerRef        string           `xml:"perRef"`
	IdeADC        []*IdeADC        `xml:"ideADC"`
	DetInfoPerRef []*DetInfoPerRef `xml:"detInfoPerRef"`
}
