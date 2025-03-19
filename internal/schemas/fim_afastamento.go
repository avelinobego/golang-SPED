package schemas

import (
	"encoding/xml"
)

// FimAfastamento ...
type FimAfastamento struct {
	XMLName     xml.Name `xml:"fimAfastamento"`
	DtTermAfast string   `xml:"dtTermAfast"`
}
