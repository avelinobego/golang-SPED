package schemas

import (
	"encoding/xml"
)

// InfoCCP ...
type InfoCCP struct {
	XMLName xml.Name `xml:"infoCCP"`
	DtCCP   string   `xml:"dtCCP"`
	TpCCP   int8     `xml:"tpCCP"`
	CnpjCCP string   `xml:"cnpjCCP"`
}
