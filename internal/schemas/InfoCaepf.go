package schemas

import (
	"encoding/xml"
)

// InfoCaepf ...
type InfoCaepf struct {
	XMLName xml.Name `xml:"infoCaepf"`
	TpCaepf int8     `xml:"tpCaepf"`
}
