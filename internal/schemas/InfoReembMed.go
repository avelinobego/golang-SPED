package schemas

import (
	"encoding/xml"
)

// InfoReembMed ...
type InfoReembMed struct {
	XMLName      xml.Name        `xml:"infoReembMed"`
	IndOrgReemb  string          `xml:"indOrgReemb"`
	CnpjOper     string          `xml:"cnpjOper"`
	RegANS       string          `xml:"regANS"`
	DetReembTit  []string        `xml:"detReembTit"`
	InfoReembDep []*InfoReembDep `xml:"infoReembDep"`
}
