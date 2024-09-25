package schemes

import (
	"encoding/xml"
)

// InfoPgtoExt ...
type InfoPgtoExt struct {
	XMLName   xml.Name `xml:"infoPgtoExt"`
	IndNIF    string   `xml:"indNIF"`
	NifBenef  string   `xml:"nifBenef"`
	FrmTribut string   `xml:"frmTribut"`
	EndExt    *EndExt  `xml:"endExt"`
}
