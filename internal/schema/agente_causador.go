package schema

import (
	"encoding/xml"
)

// AgenteCausador ...
type AgenteCausador struct {
	XMLName         xml.Name `xml:"agenteCausador"`
	CodAgntCausador int      `xml:"codAgntCausador"`
}
