package schemes

import (
	"encoding/xml"
)

// InfoTributos ...
type InfoTributos struct {
	XMLName       xml.Name         `xml:"infoTributos"`
	PerRef        string           `xml:"perRef"`
	InfoCRContrib []*InfoCRContrib `xml:"infoCRContrib"`
}
