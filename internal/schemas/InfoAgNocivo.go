package schemas

import (
	"encoding/xml"
)

// InfoAgNocivo ...
type InfoAgNocivo struct {
	XMLName xml.Name `xml:"infoAgNocivo"`
	GrauExp string   `xml:"grauExp"`
}
