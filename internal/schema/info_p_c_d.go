package schema

import (
	"encoding/xml"
)

// InfoPCD ...
type InfoPCD struct {
	XMLName   xml.Name `xml:"infoPCD"`
	NrProcJud string   `xml:"nrProcJud"`
}
