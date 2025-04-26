package events

import "encoding/xml"

func MakeESocial(versao string, evento any) any {
	result := struct {
		Evento  any
		XMLName xml.Name
	}{
		XMLName: xml.Name{
			Local: "eSocial",
			Space: versao,
		},
		Evento: evento,
	}
	return result
}
