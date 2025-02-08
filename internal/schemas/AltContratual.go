package schemas

import (
	"encoding/xml"
)

// AltContratual ...
type AltContratual struct {
	XMLName     xml.Name `xml:"altContratual"`
	DtAlteracao string   `xml:"dtAlteracao"`
	DtEf        string   `xml:"dtEf"`
	DscAlt      string   `xml:"dscAlt"`
	Vinculo     *Vinculo `xml:"vinculo"`
}
