package schemas

import (
	"encoding/xml"
)

// TdadosProc is Dados do processo.
type TdadosProc struct {
	XMLName      xml.Name      `xml:"T_dadosProc"`
	IndAutoria   int8          `xml:"indAutoria"`
	IndMatProc   int8          `xml:"indMatProc"`
	Observacao   string        `xml:"observacao"`
	DadosProcJud *DadosProcJud `xml:"dadosProcJud"`
	InfoSusp     []*InfoSusp   `xml:"infoSusp"`
}
