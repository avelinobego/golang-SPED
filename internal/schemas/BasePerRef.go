package schemas

import (
	"encoding/xml"
)

// BasePerRef ...
type BasePerRef struct {
	XMLName         xml.Name `xml:"basePerRef"`
	PerRef          string   `xml:"perRef"`
	CodCateg        string   `xml:"codCateg"`
	TpValorProcTrab int8     `xml:"tpValorProcTrab"`
	RemFGTSProcTrab string   `xml:"remFGTSProcTrab"`
	DpsFGTSProcTrab string   `xml:"dpsFGTSProcTrab"`
	RemFGTSSefip    string   `xml:"remFGTSSefip"`
	DpsFGTSSefip    string   `xml:"dpsFGTSSefip"`
	RemFGTSDecAnt   string   `xml:"remFGTSDecAnt"`
	DpsFGTSDecAnt   string   `xml:"dpsFGTSDecAnt"`
}
