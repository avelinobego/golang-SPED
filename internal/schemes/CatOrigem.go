package schemes

import (
	"encoding/xml"
)

// CatOrigem ...
type CatOrigem struct {
	XMLName      xml.Name `xml:"catOrigem"`
	NrRecCatOrig string   `xml:"nrRecCatOrig"`
}
