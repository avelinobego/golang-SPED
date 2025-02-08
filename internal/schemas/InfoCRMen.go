package schemas

import (
	"encoding/xml"
)

// InfoCRMen ...
type InfoCRMen struct {
	XMLName xml.Name `xml:"infoCRMen"`
	CRMen   string   `xml:"CRMen"`
	VrCRMen string   `xml:"vrCRMen"`
}
