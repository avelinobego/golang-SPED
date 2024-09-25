package schemes

import (
	"encoding/xml"
)

// IdeAdquir ...
type IdeAdquir struct {
	XMLName  xml.Name `xml:"ideAdquir"`
	TpInsc   string   `xml:"tpInsc"`
	NrInsc   string   `xml:"nrInsc"`
	VrComerc string   `xml:"vrComerc"`
	Nfs      []*Nfs   `xml:"nfs"`
}
