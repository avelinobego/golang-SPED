package schemes

import (
	"encoding/xml"
)

// TitensRemunrppsperApur ...
type TitensRemunrppsperApur struct {
	XMLName   xml.Name `xml:"T_itensRemun_rpps_perApur"`
	DescFolha string   `xml:"descFolha"`
}
