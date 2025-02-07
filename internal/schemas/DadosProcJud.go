package schemes

import (
	"encoding/xml"
)

// DadosProcJud ...
type DadosProcJud struct {
	XMLName  xml.Name `xml:"dadosProcJud"`
	UfVara   string   `xml:"ufVara"`
	CodMunic string   `xml:"codMunic"`
	IdVara   int      `xml:"idVara"`
}
