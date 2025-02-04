package schemes

import (
	"encoding/xml"
)

// ESocial ...
type ESocial[E any] struct {
	XMLName   xml.Name `xml:"eSocial"`
	Namespace string   `xml:"xmlns,attr" validate:"required"`
	Evento    *E       `validate:"required"`
}
