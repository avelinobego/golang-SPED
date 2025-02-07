package schemes

import (
	"encoding/xml"
)

// InfoEmprParcial ...
type InfoEmprParcial struct {
	XMLName       xml.Name `xml:"infoEmprParcial"`
	TpInscContrat string   `xml:"tpInscContrat"`
	NrInscContrat string   `xml:"nrInscContrat"`
	TpInscProp    string   `xml:"tpInscProp"`
	NrInscProp    string   `xml:"nrInscProp"`
}
