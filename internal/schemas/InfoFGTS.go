package schemes

import (
	"encoding/xml"
)

// InfoFGTS ...
type InfoFGTS struct {
	XMLName          xml.Name `xml:"infoFGTS"`
	VrBcFGTSProcTrab string   `xml:"vrBcFGTSProcTrab"`
	VrBcFGTSSefip    string   `xml:"vrBcFGTSSefip"`
	VrBcFGTSDecAnt   string   `xml:"vrBcFGTSDecAnt"`
}
