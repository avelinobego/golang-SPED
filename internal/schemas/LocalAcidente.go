package schemes

import (
	"encoding/xml"
)

// LocalAcidente ...
type LocalAcidente struct {
	XMLName      xml.Name      `xml:"localAcidente"`
	TpLocal      int8          `xml:"tpLocal"`
	DscLocal     string        `xml:"dscLocal"`
	TpLograd     string        `xml:"tpLograd"`
	DscLograd    string        `xml:"dscLograd"`
	NrLograd     string        `xml:"nrLograd"`
	Complemento  string        `xml:"complemento"`
	Bairro       string        `xml:"bairro"`
	Cep          string        `xml:"cep"`
	CodMunic     string        `xml:"codMunic"`
	Uf           string        `xml:"uf"`
	Pais         string        `xml:"pais"`
	CodPostal    string        `xml:"codPostal"`
	IdeLocalAcid *IdeLocalAcid `xml:"ideLocalAcid"`
}
