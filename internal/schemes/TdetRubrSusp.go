package schemes

import (
	"encoding/xml"
)

// TdetRubrSusp ...
type TdetRubrSusp struct {
	XMLName         xml.Name           `xml:"T_detRubrSusp"`
	CodRubr         string             `xml:"codRubr"`
	IdeTabRubr      string             `xml:"ideTabRubr"`
	VrRubr          string             `xml:"vrRubr"`
	IdeProcessoFGTS []*IdeProcessoFGTS `xml:"ideProcessoFGTS"`
}
