package schema

import (
	"encoding/xml"
)

// BasesAquis ...
type BasesAquis struct {
	XMLName     xml.Name `xml:"basesAquis"`
	IndAquis    int8     `xml:"indAquis"`
	VlrAquis    string   `xml:"vlrAquis"`
	VrCPDescPR  string   `xml:"vrCPDescPR"`
	VrCPNRet    string   `xml:"vrCPNRet"`
	VrRatNRet   string   `xml:"vrRatNRet"`
	VrSenarNRet string   `xml:"vrSenarNRet"`
	VrCPCalcPR  string   `xml:"vrCPCalcPR"`
	VrRatDescPR string   `xml:"vrRatDescPR"`
	VrRatCalcPR string   `xml:"vrRatCalcPR"`
	VrSenarDesc string   `xml:"vrSenarDesc"`
	VrSenarCalc string   `xml:"vrSenarCalc"`
}
