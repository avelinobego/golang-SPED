package schemes

import (
	"encoding/xml"
)

// InfoAtiv ...
type InfoAtiv struct {
	XMLName    xml.Name `xml:"infoAtiv"`
	DscAtivDes string   `xml:"dscAtivDes"`
}
