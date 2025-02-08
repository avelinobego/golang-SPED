package schemas

import (
	"encoding/xml"
)

// RespMonit ...
type RespMonit struct {
	XMLName xml.Name `xml:"respMonit"`
	CpfResp string   `xml:"cpfResp"`
	NmResp  string   `xml:"nmResp"`
	NrCRM   string   `xml:"nrCRM"`
	UfCRM   string   `xml:"ufCRM"`
}
