package schemes

import (
	"encoding/xml"
)

// InfoDeficiencia ...
type InfoDeficiencia struct {
	XMLName        xml.Name `xml:"infoDeficiencia"`
	DefFisica      string   `xml:"defFisica"`
	DefVisual      string   `xml:"defVisual"`
	DefAuditiva    string   `xml:"defAuditiva"`
	DefMental      string   `xml:"defMental"`
	DefIntelectual string   `xml:"defIntelectual"`
	ReabReadap     string   `xml:"reabReadap"`
	Observacao     string   `xml:"observacao"`
}
