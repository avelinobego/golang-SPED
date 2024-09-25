package schemes

import (
	"encoding/xml"
)

// PrevidCompl ...
type PrevidCompl struct {
	XMLName         xml.Name `xml:"previdCompl"`
	TpPrev          string   `xml:"tpPrev"`
	CnpjEntidPC     string   `xml:"cnpjEntidPC"`
	VlrDedPC        string   `xml:"vlrDedPC"`
	VlrDedPC13      string   `xml:"vlrDedPC13"`
	VlrPatrocFunp   string   `xml:"vlrPatrocFunp"`
	VlrPatrocFunp13 string   `xml:"vlrPatrocFunp13"`
}
