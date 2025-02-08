package schemas

import (
	"encoding/xml"
)

// LocalTrabalho ...
type LocalTrabalho struct {
	XMLName        xml.Name `xml:"localTrabalho"`
	LocalTrabGeral string   `xml:"localTrabGeral"`
	LocalTempDom   string   `xml:"localTempDom"`
}
