package schemes

import (
	"encoding/xml"
)

// Abono ...
type Abono struct {
	XMLName xml.Name `xml:"abono"`
	AnoBase int      `xml:"anoBase"`
}
