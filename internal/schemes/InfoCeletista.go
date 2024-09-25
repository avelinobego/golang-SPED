package schemes

import (
	"encoding/xml"
)

// InfoCeletista ...
type InfoCeletista struct {
	XMLName           xml.Name        `xml:"infoCeletista"`
	TpRegJor          string          `xml:"tpRegJor"`
	NatAtividade      string          `xml:"natAtividade"`
	DtBase            string          `xml:"dtBase"`
	CnpjSindCategProf string          `xml:"cnpjSindCategProf"`
	TrabTemporario    *TrabTemporario `xml:"trabTemporario"`
	Aprend            string          `xml:"aprend"`
}
