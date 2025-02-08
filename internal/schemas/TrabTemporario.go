package schemas

import (
	"encoding/xml"
)

// TrabTemporario ...
type TrabTemporario struct {
	XMLName   xml.Name `xml:"trabTemporario"`
	JustProrr string   `xml:"justProrr"`
}
