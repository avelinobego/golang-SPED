package schemas

import (
	"encoding/xml"
)

// IdeEvento ...
type IdeEvento struct {
	XMLName      xml.Name `xml:"ideEvento"`
	NrRecArqBase string   `xml:"nrRecArqBase"`
}
