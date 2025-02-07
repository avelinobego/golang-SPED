package schemes

import (
	"encoding/xml"
)

// TrabImig ...
type TrabImig struct {
	XMLName  xml.Name `xml:"trabImig"`
	TmpResid string   `xml:"tmpResid"`
	CondIng  string   `xml:"condIng"`
}
