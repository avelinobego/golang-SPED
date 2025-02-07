package schemes

import (
	"encoding/xml"
)

// InfoMV ...
type InfoMV struct {
	XMLName       xml.Name         `xml:"infoMV"`
	IndMV         string           `xml:"indMV"`
	RemunOutrEmpr []*RemunOutrEmpr `xml:"remunOutrEmpr"`
}
