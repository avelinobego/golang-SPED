package schemas

import (
	"encoding/xml"
)

// EpcEpi ...
type EpcEpi struct {
	XMLName   xml.Name  `xml:"epcEpi"`
	UtilizEPC int8      `xml:"utilizEPC"`
	EficEpc   string    `xml:"eficEpc"`
	UtilizEPI int8      `xml:"utilizEPI"`
	EficEpi   string    `xml:"eficEpi"`
	Epi       []*Epi    `xml:"epi"`
	EpiCompl  *EpiCompl `xml:"epiCompl"`
}
