package schema

import (
	"encoding/xml"
)

// InfoPerAnt ...
type InfoPerAnt struct {
	XMLName     xml.Name      `xml:"infoPerAnt"`
	RemunOrgSuc string        `xml:"remunOrgSuc"`
	IdePeriodo  []*IdePeriodo `xml:"idePeriodo"`
}
