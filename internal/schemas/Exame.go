package schemes

import (
	"encoding/xml"
)

// Exame ...
type Exame struct {
	XMLName       xml.Name `xml:"exame"`
	DtExm         string   `xml:"dtExm"`
	ProcRealizado int      `xml:"procRealizado"`
	ObsProc       string   `xml:"obsProc"`
	OrdExame      int8     `xml:"ordExame"`
	IndResult     int8     `xml:"indResult"`
}
