package schemes

import (
	"encoding/xml"
)

// BasePerApur ...
type BasePerApur struct {
	XMLName  xml.Name `xml:"basePerApur"`
	TpValor  int8     `xml:"tpValor"`
	IndIncid string   `xml:"indIncid"`
	BaseFGTS string   `xml:"baseFGTS"`
	VrFGTS   string   `xml:"vrFGTS"`
}
