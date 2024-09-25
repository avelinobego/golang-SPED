package schemes

import (
	"encoding/xml"
)

// DadosCompl ...
type DadosCompl struct {
	XMLName     xml.Name     `xml:"dadosCompl"`
	InfoProcJud *InfoProcJud `xml:"infoProcJud"`
	InfoCCP     *InfoCCP     `xml:"infoCCP"`
}
