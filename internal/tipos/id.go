package tipos

import "encoding/xml"

type Id struct {
}

func (i Id) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	//TODO: criar os mecanismos do ID aqui
	return xml.Attr{Name: name, Value: "Oi"}, nil
}
