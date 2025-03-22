package schema

import (
	"encoding/xml"
)

// InfoReativ ...
type InfoReativ struct {
	XMLName      xml.Name `xml:"infoReativ"`
	DtEfetReativ string   `xml:"dtEfetReativ"`
	DtEfeito     string   `xml:"dtEfeito"`
}
