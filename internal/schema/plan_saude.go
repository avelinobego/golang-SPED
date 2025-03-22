package schema

import (
	"encoding/xml"
)

// PlanSaude ...
type PlanSaude struct {
	XMLName     xml.Name      `xml:"planSaude"`
	CnpjOper    string        `xml:"cnpjOper"`
	RegANS      string        `xml:"regANS"`
	VlrSaudeTit string        `xml:"vlrSaudeTit"`
	InfoDepSau  []*InfoDepSau `xml:"infoDepSau"`
}
