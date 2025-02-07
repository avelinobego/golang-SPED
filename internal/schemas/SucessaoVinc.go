package schemes

import (
	"encoding/xml"
)

// SucessaoVinc ...
type SucessaoVinc struct {
	XMLName      xml.Name `xml:"sucessaoVinc"`
	CnpjOrgaoAnt string   `xml:"cnpjOrgaoAnt"`
	MatricAnt    string   `xml:"matricAnt"`
	DtExercicio  string   `xml:"dtExercicio"`
	Observacao   string   `xml:"observacao"`
}
