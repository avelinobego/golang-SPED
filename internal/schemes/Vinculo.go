package schemes

import (
	"encoding/xml"
)

// Vinculo ...
type Vinculo struct {
	XMLName        xml.Name        `xml:"vinculo"`
	TpRegPrev      string          `xml:"tpRegPrev"`
	InfoRegimeTrab *InfoRegimeTrab `xml:"infoRegimeTrab"`
	InfoContrato   *InfoContrato   `xml:"infoContrato"`
}
