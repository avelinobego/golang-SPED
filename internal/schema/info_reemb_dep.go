package schema

import (
	"encoding/xml"
)

// InfoReembDep ...
type InfoReembDep struct {
	XMLName     xml.Name `xml:"infoReembDep"`
	CpfBenef    string   `xml:"cpfBenef"`
	DetReembDep []string `xml:"detReembDep"`
}
