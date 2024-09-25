package schemes

import (
	"encoding/xml"
)

// InfoExpRisco ...
type InfoExpRisco struct {
	XMLName       xml.Name   `xml:"infoExpRisco"`
	DtIniCondicao string     `xml:"dtIniCondicao"`
	DtFimCondicao string     `xml:"dtFimCondicao"`
	InfoAmb       []*InfoAmb `xml:"infoAmb"`
	InfoAtiv      *InfoAtiv  `xml:"infoAtiv"`
	AgNoc         []*AgNoc   `xml:"agNoc"`
	RespReg       []*RespReg `xml:"respReg"`
	Obs           *Obs       `xml:"obs"`
}
