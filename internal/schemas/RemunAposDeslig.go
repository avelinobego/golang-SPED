package schemes

import (
	"encoding/xml"
)

// RemunAposDeslig ...
type RemunAposDeslig struct {
	XMLName    xml.Name `xml:"remunAposDeslig"`
	IndRemun   int8     `xml:"indRemun"`
	DtFimRemun string   `xml:"dtFimRemun"`
}
