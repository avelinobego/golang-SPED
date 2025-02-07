package schemes

import (
	"encoding/xml"
)

// InfoOrgInternacional ...
type InfoOrgInternacional struct {
	XMLName            xml.Name `xml:"infoOrgInternacional"`
	IndAcordoIsenMulta int8     `xml:"indAcordoIsenMulta"`
}
