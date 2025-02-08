package schemas

import (
	"encoding/xml"
)

// InfoMandElet ...
type InfoMandElet struct {
	XMLName       xml.Name `xml:"infoMandElet"`
	CategOrig     string   `xml:"categOrig"`
	CnpjOrig      string   `xml:"cnpjOrig"`
	MatricOrig    string   `xml:"matricOrig"`
	DtExercOrig   string   `xml:"dtExercOrig"`
	IndRemunCargo string   `xml:"indRemunCargo"`
	TpRegTrab     string   `xml:"tpRegTrab"`
	TpRegPrev     string   `xml:"tpRegPrev"`
}
