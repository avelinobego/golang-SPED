package schemas

import (
	"encoding/xml"
)

// InfoCS ...
type InfoCS struct {
	XMLName       xml.Name         `xml:"infoCS"`
	NrRecArqBase  string           `xml:"nrRecArqBase"`
	IndExistInfo  int8             `xml:"indExistInfo"`
	InfoCPSeg     *InfoCPSeg       `xml:"infoCPSeg"`
	InfoContrib   *InfoContrib     `xml:"infoContrib"`
	IdeEstab      []*IdeEstab      `xml:"ideEstab"`
	InfoCRContrib []*InfoCRContrib `xml:"infoCRContrib"`
}
