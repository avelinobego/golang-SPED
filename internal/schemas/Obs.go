package schemas

import (
	"encoding/xml"
)

// Obs ...
type Obs struct {
	XMLName  xml.Name `xml:"obs"`
	ObsCompl string   `xml:"obsCompl"`
}
