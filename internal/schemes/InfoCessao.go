package schemes

import (
	"encoding/xml"
)

// InfoCessao ...
type InfoCessao struct {
	XMLName   xml.Name   `xml:"infoCessao"`
	IniCessao *IniCessao `xml:"iniCessao"`
	FimCessao *FimCessao `xml:"fimCessao"`
}
