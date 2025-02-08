package schemas

import (
	"encoding/xml"
)

// CalcTerc ...
type CalcTerc struct {
	XMLName     xml.Name `xml:"calcTerc"`
	TpCR        int      `xml:"tpCR"`
	VrCsSegTerc string   `xml:"vrCsSegTerc"`
	VrDescTerc  string   `xml:"vrDescTerc"`
}
