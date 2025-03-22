package schema

import (
	"encoding/xml"
)

// ESocial ...
type ESocial[evento any] struct {
	XMLName   xml.Name `xml:"eSocial"`
	Namespace string   `xml:"xmlns,attr" validate:"required"`
	Evento    *evento  `validate:"required"`
}
