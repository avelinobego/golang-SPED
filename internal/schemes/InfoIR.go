package schemes

import (
	"encoding/xml"
)

// InfoIR ...
type InfoIR struct {
	XMLName        xml.Name          `xml:"infoIR"`
	TpInfoIR       int               `xml:"tpInfoIR"`
	Valor          string            `xml:"valor"`
	DescRendimento string            `xml:"descRendimento"`
	InfoProcJudRub []*InfoProcJudRub `xml:"infoProcJudRub"`
}
