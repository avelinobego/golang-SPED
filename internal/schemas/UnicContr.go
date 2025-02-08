package schemas

import (
	"encoding/xml"
)

// UnicContr ...
type UnicContr struct {
	XMLName  xml.Name `xml:"unicContr"`
	MatUnic  string   `xml:"matUnic"`
	CodCateg string   `xml:"codCateg"`
	DtInicio string   `xml:"dtInicio"`
}
