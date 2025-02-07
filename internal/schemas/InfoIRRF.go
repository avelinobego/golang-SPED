package schemes

import (
	"encoding/xml"
)

// InfoIRRF ...
type InfoIRRF struct {
	XMLName      xml.Name     `xml:"infoIRRF"`
	NrRecArqBase string       `xml:"nrRecArqBase"`
	IndExistInfo int8         `xml:"indExistInfo"`
	InfoCRMen    []*InfoCRMen `xml:"infoCRMen"`
	InfoCRDia    []*InfoCRDia `xml:"infoCRDia"`
}
