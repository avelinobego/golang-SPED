package schemes

import (
	"encoding/xml"
)

// BaseCalculo ...
type BaseCalculo struct {
	XMLName      xml.Name      `xml:"baseCalculo"`
	VrBcCpMensal string        `xml:"vrBcCpMensal"`
	VrBcCp13     string        `xml:"vrBcCp13"`
	InfoAgNocivo *InfoAgNocivo `xml:"infoAgNocivo"`
}
