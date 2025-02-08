package schemas

import (
	"encoding/xml"
)

// TinfoAgNocivo ...
type TinfoAgNocivo struct {
	XMLName xml.Name `xml:"T_infoAgNocivo"`
	GrauExp string   `xml:"grauExp"`
}
