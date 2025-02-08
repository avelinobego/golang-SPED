package schemas

import (
	"encoding/xml"
)

// Cat ...
type Cat struct {
	XMLName          xml.Name        `xml:"cat"`
	DtAcid           string          `xml:"dtAcid"`
	TpAcid           int8            `xml:"tpAcid"`
	HrAcid           string          `xml:"hrAcid"`
	HrsTrabAntesAcid string          `xml:"hrsTrabAntesAcid"`
	TpCat            int8            `xml:"tpCat"`
	IndCatObito      string          `xml:"indCatObito"`
	DtObito          string          `xml:"dtObito"`
	IndComunPolicia  string          `xml:"indComunPolicia"`
	CodSitGeradora   int             `xml:"codSitGeradora"`
	IniciatCAT       int8            `xml:"iniciatCAT"`
	ObsCAT           string          `xml:"obsCAT"`
	UltDiaTrab       string          `xml:"ultDiaTrab"`
	HouveAfast       string          `xml:"houveAfast"`
	LocalAcidente    *LocalAcidente  `xml:"localAcidente"`
	ParteAtingida    *ParteAtingida  `xml:"parteAtingida"`
	AgenteCausador   *AgenteCausador `xml:"agenteCausador"`
	Atestado         *Atestado       `xml:"atestado"`
	CatOrigem        *CatOrigem      `xml:"catOrigem"`
}
