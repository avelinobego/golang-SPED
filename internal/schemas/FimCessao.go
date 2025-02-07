package schemes

import (
	"encoding/xml"
)

// FimCessao ...
type FimCessao struct {
	XMLName      xml.Name `xml:"fimCessao"`
	DtTermCessao string   `xml:"dtTermCessao"`
}
