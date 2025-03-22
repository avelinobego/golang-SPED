package schema

import (
	"encoding/xml"
)

// DespProcJud ...
type DespProcJud struct {
	XMLName          xml.Name `xml:"despProcJud"`
	VlrDespCustas    string   `xml:"vlrDespCustas"`
	VlrDespAdvogados string   `xml:"vlrDespAdvogados"`
}
