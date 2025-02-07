package schemes

import (
	"encoding/xml"
)

// TpComerc ...
type TpComerc struct {
	XMLName     xml.Name       `xml:"tpComerc"`
	IndComerc   string         `xml:"indComerc"`
	VrTotCom    string         `xml:"vrTotCom"`
	IdeAdquir   []*IdeAdquir   `xml:"ideAdquir"`
	InfoProcJud []*InfoProcJud `xml:"infoProcJud"`
}
