package schemes

import (
	"encoding/xml"
)

// TideProcesso is c) Se {tpProc}(./tpProc) = [4], deve possuir 16 (dezesseis) algarismos.
type TideProcesso struct {
	XMLName  xml.Name `xml:"T_ideProcesso"`
	TpProc   string   `xml:"tpProc"`
	NrProc   string   `xml:"nrProc"`
	IniValid string   `xml:"iniValid"`
	FimValid string   `xml:"fimValid"`
}
