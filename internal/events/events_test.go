package events_test

import (
	"encoding/xml"
	"sped/internal/events"
	"testing"
)

func TestESocial(t *testing.T) {
	e := events.MakeESocial("xpto", events.EvtAdmissao{
		Nome: "Avelino Bego",
	})

	if bs, e := xml.Marshal(e); e == nil {
		println(string(bs))
	} else {
		t.Fatal(e)
	}
}
