package schema

import (
	"encoding/xml"
)

// Toxicologico ...
type Toxicologico struct {
	XMLName     xml.Name `xml:"toxicologico"`
	DtExame     string   `xml:"dtExame"`
	CnpjLab     string   `xml:"cnpjLab"`
	CodSeqExame string   `xml:"codSeqExame"`
	NmMed       string   `xml:"nmMed"`
	NrCRM       string   `xml:"nrCRM"`
	UfCRM       string   `xml:"ufCRM"`
}
