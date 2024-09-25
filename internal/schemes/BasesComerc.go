package schemes

import (
	"encoding/xml"
)

// BasesComerc ...
type BasesComerc struct {
	XMLName     xml.Name `xml:"basesComerc"`
	IndComerc   string   `xml:"indComerc"`
	VrBcComPR   string   `xml:"vrBcComPR"`
	VrCPSusp    string   `xml:"vrCPSusp"`
	VrRatSusp   string   `xml:"vrRatSusp"`
	VrSenarSusp string   `xml:"vrSenarSusp"`
}
