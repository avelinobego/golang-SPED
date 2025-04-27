package tipos

type Trabalhador struct {
	XMLName    struct{}       `xml:"trabalhador"`
	CpfTrab    uint32         `xml:"cpfTrab"`
	NmTrab     string         `xml:"nmTrab"`
	Sexo       string         `xml:"sexo"`
	RacaCor    uint32         `xml:"racaCor"`
	EstCiv     uint32         `xml:"estCiv,omitempty"`
	GrauInstr  string         `xml:"grauInstr"`
	NmSoc      string         `xml:"nmSoc"`
	Nascimento TipoNascimento `xml:"nascimento"`
	
}
