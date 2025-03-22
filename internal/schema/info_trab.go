package schema

import (
	"encoding/xml"
)

// InfoTrab ...
type InfoTrab struct {
	XMLName xml.Name `xml:"infoTrab"`
	InfoApr *InfoApr `xml:"infoApr"`
	InfoPCD *InfoPCD `xml:"infoPCD"`
}
