package schemes

import (
	"encoding/xml"
)

// RemunAposTerm ...
type RemunAposTerm struct {
	XMLName    xml.Name `xml:"remunAposTerm"`
	IndRemun   int8     `xml:"indRemun"`
	DtFimRemun string   `xml:"dtFimRemun"`
}
