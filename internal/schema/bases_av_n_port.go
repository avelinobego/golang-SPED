package schema

import (
	"encoding/xml"
)

// BasesAvNPort ...
type BasesAvNPort struct {
	XMLName  xml.Name `xml:"basesAvNPort"`
	VrBcCp00 string   `xml:"vrBcCp00"`
	VrBcCp15 string   `xml:"vrBcCp15"`
	VrBcCp20 string   `xml:"vrBcCp20"`
	VrBcCp25 string   `xml:"vrBcCp25"`
	VrBcCp13 string   `xml:"vrBcCp13"`
	VrDescCP string   `xml:"vrDescCP"`
}
