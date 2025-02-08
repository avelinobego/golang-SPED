package schemas

import (
	"encoding/xml"
)

// RemunOutrEmpr ...
type RemunOutrEmpr struct {
	XMLName    xml.Name `xml:"remunOutrEmpr"`
	TpInsc     string   `xml:"tpInsc"`
	NrInsc     string   `xml:"nrInsc"`
	CodCateg   string   `xml:"codCateg"`
	VlrRemunOE string   `xml:"vlrRemunOE"`
}
