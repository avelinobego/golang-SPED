package tipos

type TipoNascimento struct {
	XMLName    struct{} `xml:"nascimento"`
	DtNascto   DataNasc `xml:"dtNascto"`
	PaisNascto string   `xml:"paisNascto"`
}
