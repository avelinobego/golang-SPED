package schema

import (
	"encoding/xml"
)

// TotInfoIR ...
type TotInfoIR struct {
	XMLName         xml.Name           `xml:"totInfoIR"`
	ConsolidApurMen []*ConsolidApurMen `xml:"consolidApurMen"`
}
