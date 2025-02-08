package schemas

import (
	"encoding/xml"
)

// ESocial ...
type ESocial struct {
	XMLName   xml.Name `xml:"eSocial"`
	Namespace string   `xml:"xmlns,attr" validate:"required"`
	Evento    *any     `validate:"required"`
}
