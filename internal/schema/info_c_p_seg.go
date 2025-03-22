package schema

import (
	"encoding/xml"
)

// InfoCPSeg ...
type InfoCPSeg struct {
	XMLName  xml.Name `xml:"infoCPSeg"`
	VrDescCP string   `xml:"vrDescCP"`
	VrCpSeg  string   `xml:"vrCpSeg"`
}
