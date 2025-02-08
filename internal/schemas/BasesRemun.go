package schemas

import (
	"encoding/xml"
)

// BasesRemun ...
type BasesRemun struct {
	XMLName  xml.Name `xml:"basesRemun"`
	IndIncid int8     `xml:"indIncid"`
	CodCateg string   `xml:"codCateg"`
	BasesCp  *BasesCp `xml:"basesCp"`
}
