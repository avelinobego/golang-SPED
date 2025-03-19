package schemas

import (
	"encoding/xml"
)

// InfoRetif ...
type InfoRetif struct {
	XMLName   xml.Name `xml:"infoRetif"`
	OrigRetif int8     `xml:"origRetif"`
	TpProc    int8     `xml:"tpProc"`
	NrProc    string   `xml:"nrProc"`
}
