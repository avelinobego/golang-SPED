package schemas

import (
	"encoding/xml"
)

// TotApurMen ...
type TotApurMen struct {
	XMLName            xml.Name `xml:"totApurMen"`
	CRMen              string   `xml:"CRMen"`
	VlrRendTrib        string   `xml:"vlrRendTrib"`
	VlrRendTrib13      string   `xml:"vlrRendTrib13"`
	VlrPrevOficial     string   `xml:"vlrPrevOficial"`
	VlrPrevOficial13   string   `xml:"vlrPrevOficial13"`
	VlrCRMen           string   `xml:"vlrCRMen"`
	VlrCR13Men         string   `xml:"vlrCR13Men"`
	VlrParcIsenta65    string   `xml:"vlrParcIsenta65"`
	VlrParcIsenta65Dec string   `xml:"vlrParcIsenta65Dec"`
	VlrDiarias         string   `xml:"vlrDiarias"`
	VlrAjudaCusto      string   `xml:"vlrAjudaCusto"`
	VlrIndResContrato  string   `xml:"vlrIndResContrato"`
	VlrAbonoPec        string   `xml:"vlrAbonoPec"`
	VlrRendMoleGrave   string   `xml:"vlrRendMoleGrave"`
	VlrRendMoleGrave13 string   `xml:"vlrRendMoleGrave13"`
	VlrAuxMoradia      string   `xml:"vlrAuxMoradia"`
	VlrBolsaMedico     string   `xml:"vlrBolsaMedico"`
	VlrBolsaMedico13   string   `xml:"vlrBolsaMedico13"`
	VlrJurosMora       string   `xml:"vlrJurosMora"`
	VlrIsenOutros      string   `xml:"vlrIsenOutros"`
	DescRendimento     string   `xml:"descRendimento"`
}
