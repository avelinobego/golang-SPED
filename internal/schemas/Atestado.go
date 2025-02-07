package schemes

import (
	"encoding/xml"
)

// Atestado ...
type Atestado struct {
	XMLName       xml.Name  `xml:"atestado"`
	DtAtendimento string    `xml:"dtAtendimento"`
	HrAtendimento string    `xml:"hrAtendimento"`
	IndInternacao string    `xml:"indInternacao"`
	DurTrat       int       `xml:"durTrat"`
	IndAfast      string    `xml:"indAfast"`
	DscLesao      int       `xml:"dscLesao"`
	DscCompLesao  string    `xml:"dscCompLesao"`
	DiagProvavel  string    `xml:"diagProvavel"`
	CodCID        string    `xml:"codCID"`
	Observacao    string    `xml:"observacao"`
	Emitente      *Emitente `xml:"emitente"`
}
