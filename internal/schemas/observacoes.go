package schemas

import (
	"encoding/xml"
)

// Observacoes ...
type Observacoes struct {
	XMLName    xml.Name `xml:"observacoes"`
	Observacao string   `xml:"observacao"`
}
