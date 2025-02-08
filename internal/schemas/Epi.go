package schemas

import (
	"encoding/xml"
)

// Epi ...
type Epi struct {
	XMLName xml.Name `xml:"epi"`
	DocAval string   `xml:"docAval"`
}
