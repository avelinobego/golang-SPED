package schema

import (
	"encoding/xml"
)

// IniCessao ...
type IniCessao struct {
	XMLName     xml.Name `xml:"iniCessao"`
	DtIniCessao string   `xml:"dtIniCessao"`
	CnpjCess    string   `xml:"cnpjCess"`
	RespRemun   string   `xml:"respRemun"`
}
