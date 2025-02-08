package schemas

import (
	"encoding/xml"
)

// BaseMudCateg ...
type BaseMudCateg struct {
	XMLName   xml.Name `xml:"baseMudCateg"`
	CodCateg  string   `xml:"codCateg"`
	VrBcCPrev string   `xml:"vrBcCPrev"`
}
