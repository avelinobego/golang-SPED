package schemes

import (
	"encoding/xml"
)

// InfoRegimeTrab ...
type InfoRegimeTrab struct {
	XMLName         xml.Name         `xml:"infoRegimeTrab"`
	InfoCeletista   *InfoCeletista   `xml:"infoCeletista"`
	InfoEstatutario *InfoEstatutario `xml:"infoEstatutario"`
}
