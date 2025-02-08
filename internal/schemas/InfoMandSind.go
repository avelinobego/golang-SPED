package schemas

import (
	"encoding/xml"
)

// InfoMandSind ...
type InfoMandSind struct {
	XMLName      xml.Name `xml:"infoMandSind"`
	CnpjSind     string   `xml:"cnpjSind"`
	InfOnusRemun int8     `xml:"infOnusRemun"`
}
