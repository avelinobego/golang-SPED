package tipos

type IdeEventoAdmissao struct {
	XMLName  struct{} `xml:"ideEvento"`
	IndRetif uint8    `xml:"indRetif"`
	NrRecibo string   `xml:"nrRecibo,omitempty"`
	TpAmb    uint8    `xml:"tpAmb"`
	ProcEmi  uint8    `xml:"procEmi"`
	VerProc  string   `xml:"verProc"`
}
