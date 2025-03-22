package schema

import (
	"encoding/xml"
)

// InfoAtivConcom ...
type InfoAtivConcom struct {
	XMLName  xml.Name `xml:"infoAtivConcom"`
	FatorMes string   `xml:"fatorMes"`
	Fator13  string   `xml:"fator13"`
}
