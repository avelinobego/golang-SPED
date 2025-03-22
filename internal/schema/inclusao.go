package schema

import (
	"encoding/xml"
)

// Inclusao ...
type Inclusao struct {
	XMLName      xml.Name `xml:"inclusao"`
	IdeRubrica   string   `xml:"ideRubrica"`
	DadosRubrica string   `xml:"dadosRubrica"`
}
