package schemes

import (
	"encoding/xml"
)

// Alteracao ...
type Alteracao struct {
	XMLName      xml.Name `xml:"alteracao"`
	IdeRubrica   string   `xml:"ideRubrica"`
	DadosRubrica string   `xml:"dadosRubrica"`
	NovaValidade string   `xml:"novaValidade"`
}
