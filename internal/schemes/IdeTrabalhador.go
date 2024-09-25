package schemes

import (
	"encoding/xml"
)

// IdeTrabalhador ...
type IdeTrabalhador struct {
	XMLName     xml.Name     `xml:"ideTrabalhador"`
	CpfTrab     string       `xml:"cpfTrab"`
	InfoComplem *InfoComplem `xml:"infoComplem"`
}
