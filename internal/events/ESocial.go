package events

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

func MakeESocial(versao string, evento any) any {

	t := reflect.TypeOf(evento)

	var event_name string
	if f, ok := t.FieldByName("XMLName"); ok {
		event_name, _ = f.Tag.Lookup("xml")
	}

	result := struct {
		Evento  any
		XMLName xml.Name
	}{
		XMLName: xml.Name{
			Local: "eSocial",
			Space: fmt.Sprintf("http://www.esocial.gov.br/schema/evt/%s/%s", event_name, versao),
		},
		Evento: evento,
	}

	return result
}
