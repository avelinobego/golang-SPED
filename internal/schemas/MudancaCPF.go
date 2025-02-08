package schemas

import (
	"encoding/xml"
)

// MudancaCPF ...
type MudancaCPF struct {
	XMLName xml.Name `xml:"mudancaCPF"`
	NovoCPF string   `xml:"novoCPF"`
}
