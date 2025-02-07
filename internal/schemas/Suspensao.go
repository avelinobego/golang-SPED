package schemes

import (
	"encoding/xml"
)

// Suspensao ...
type Suspensao struct {
	XMLName      xml.Name `xml:"suspensao"`
	MtvSuspensao string   `xml:"mtvSuspensao"`
	DscSuspensao string   `xml:"dscSuspensao"`
}
