package tipos

type IdeEmpregador struct {
	XMLName struct{} `xml:"ideEmpregador"`
	TpInsc  uint8    `xml:"tpInsc"`
	NrInsc  string   `xml:"nrInsc"`
}
