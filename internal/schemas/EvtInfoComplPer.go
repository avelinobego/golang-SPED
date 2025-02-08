package schemas

import (
	"encoding/xml"
)

// EvtInfoComplPer ...
type EvtInfoComplPer struct {
	XMLName             xml.Name               `xml:"evtInfoComplPer"`
	IdAttr              string                 `xml:"Id,attr"`
	IdeEvento           string                 `xml:"ideEvento"`
	IdeEmpregador       string                 `xml:"ideEmpregador"`
	InfoSubstPatr       *InfoSubstPatr         `xml:"infoSubstPatr"`
	InfoSubstPatrOpPort []*InfoSubstPatrOpPort `xml:"infoSubstPatrOpPort"`
	InfoAtivConcom      *InfoAtivConcom        `xml:"infoAtivConcom"`
	InfoPercTransf11096 *InfoPercTransf11096   `xml:"infoPercTransf11096"`
}
