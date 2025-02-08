package schemas

import (
	"encoding/xml"
)

// IdeProcTrab ...
type IdeProcTrab struct {
	XMLName     xml.Name `xml:"ideProcTrab"`
	NrProcTrab  string   `xml:"nrProcTrab"`
	CpfTrab     string   `xml:"cpfTrab"`
	PerApurPgto string   `xml:"perApurPgto"`
	IdeSeqProc  int      `xml:"ideSeqProc"`
}
