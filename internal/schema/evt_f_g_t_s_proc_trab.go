package schema

import (
	"encoding/xml"
)

// EvtFGTSProcTrab ...
type EvtFGTSProcTrab struct {
	XMLName        xml.Name        `xml:"evtFGTSProcTrab"`
	IdAttr         string          `xml:"Id,attr"`
	IdeEvento      *IdeEvento      `xml:"ideEvento"`
	IdeEmpregador  *IdeEmpregador  `xml:"ideEmpregador"`
	IdeProc        *IdeProc        `xml:"ideProc"`
	IdeTrabalhador *IdeTrabalhador `xml:"ideTrabalhador"`
	InfoTrabFGTS   []*InfoTrabFGTS `xml:"infoTrabFGTS"`
}
