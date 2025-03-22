package schema

import (
	"encoding/xml"
)

// TotApurDia ...
type TotApurDia struct {
	XMLName      xml.Name `xml:"totApurDia"`
	PerApurDia   string   `xml:"perApurDia"`
	CRDia        string   `xml:"CRDia"`
	FrmTribut    string   `xml:"frmTribut"`
	PaisResidExt string   `xml:"paisResidExt"`
	VlrPagoDia   string   `xml:"vlrPagoDia"`
	VlrCRDia     string   `xml:"vlrCRDia"`
}
