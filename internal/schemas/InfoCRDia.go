package schemas

import (
	"encoding/xml"
)

// InfoCRDia ...
type InfoCRDia struct {
	XMLName    xml.Name `xml:"infoCRDia"`
	PerApurDia string   `xml:"perApurDia"`
	CRDia      string   `xml:"CRDia"`
	VrCRDia    string   `xml:"vrCRDia"`
}
