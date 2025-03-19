package schemas

import (
	"encoding/xml"
)

// InfoEstatutario ...
type InfoEstatutario struct {
	XMLName      xml.Name `xml:"infoEstatutario"`
	TpPlanRP     string   `xml:"tpPlanRP"`
	IndTetoRGPS  string   `xml:"indTetoRGPS"`
	IndAbonoPerm string   `xml:"indAbonoPerm"`
}
