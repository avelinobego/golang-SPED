package schemas

import (
	"encoding/xml"
)

// Exclusao ...
type Exclusao struct {
	XMLName    xml.Name `xml:"exclusao"`
	IdeRubrica string   `xml:"ideRubrica"`
}
