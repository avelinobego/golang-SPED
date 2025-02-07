package schemes

import (
	"encoding/xml"
)

// ProcCS ...
type ProcCS struct {
	XMLName   xml.Name `xml:"procCS"`
	NrProcJud string   `xml:"nrProcJud"`
}
