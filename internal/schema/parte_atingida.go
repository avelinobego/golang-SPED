package schema

import (
	"encoding/xml"
)

// ParteAtingida ...
type ParteAtingida struct {
	XMLName       xml.Name `xml:"parteAtingida"`
	CodParteAting int      `xml:"codParteAting"`
	Lateralidade  int8     `xml:"lateralidade"`
}
