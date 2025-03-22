package schema

import (
	"encoding/xml"
)

// InfoFGTSProcTrab ...
type InfoFGTSProcTrab struct {
	XMLName   xml.Name  `xml:"infoFGTSProcTrab"`
	TotalFGTS string    `xml:"totalFGTS"`
	IdeEstab  *IdeEstab `xml:"ideEstab"`
}
