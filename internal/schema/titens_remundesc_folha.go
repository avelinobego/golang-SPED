package schema

import (
	"encoding/xml"
)

// TitensRemundescFolha ...
type TitensRemundescFolha struct {
	XMLName   xml.Name `xml:"T_itensRemun_descFolha"`
	DescFolha string   `xml:"descFolha"`
	*TitensRemun
}
