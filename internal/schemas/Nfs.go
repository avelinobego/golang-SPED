package schemas

import (
	"encoding/xml"
)

// Nfs ...
type Nfs struct {
	XMLName     xml.Name `xml:"nfs"`
	Serie       string   `xml:"serie"`
	NrDocto     string   `xml:"nrDocto"`
	DtEmisNF    string   `xml:"dtEmisNF"`
	VlrBruto    string   `xml:"vlrBruto"`
	VrCPDescPR  string   `xml:"vrCPDescPR"`
	VrRatDescPR string   `xml:"vrRatDescPR"`
	VrSenarDesc string   `xml:"vrSenarDesc"`
}
