package schemes

import (
	"encoding/xml"
)

// IdeFolhaPagto ...
type IdeFolhaPagto struct {
	XMLName     xml.Name `xml:"ideFolhaPagto"`
	IndApuracao string   `xml:"indApuracao"`
	PerApur     string   `xml:"perApur"`
}
