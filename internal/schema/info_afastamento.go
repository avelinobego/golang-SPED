package schema

import (
	"encoding/xml"
)

// InfoAfastamento ...
type InfoAfastamento struct {
	XMLName        xml.Name        `xml:"infoAfastamento"`
	IniAfastamento *IniAfastamento `xml:"iniAfastamento"`
	InfoRetif      *InfoRetif      `xml:"infoRetif"`
	FimAfastamento *FimAfastamento `xml:"fimAfastamento"`
}
