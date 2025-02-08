package schemas

import (
	"encoding/xml"
)

// TdetVerbasdescFolha ...
type TdetVerbasdescFolha struct {
	XMLName   xml.Name `xml:"T_detVerbas_descFolha"`
	DescFolha string   `xml:"descFolha"`
	*TdetVerbas
}
