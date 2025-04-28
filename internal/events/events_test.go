package events_test

import (
	"encoding/xml"
	"sped/internal/events"
	"sped/internal/tipos"
	"testing"
)

func TestESocial(t *testing.T) {
	e := events.MakeESocial("v_S_01_03_00", events.EvtAdmissao{
		IdeEvento: tipos.IdeEventoAdmissao{
			IndRetif: 1,
			//TODO: Gerar um tipo com validador
			NrRecibo: "1.d.ddddddddddddddddddd",
			TpAmb:    1,
			ProcEmi:  1,
			VerProc:  "versao",
		},
		IdeEmpregador: tipos.IdeEmpregador{
			TpInsc: 1,
			NrInsc: "nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn",
		},
		Trabalhador: tipos.Trabalhador{
			CpfTrab: tipos.Cpf("16944301890"),
		},
	})

	if bs, e := xml.Marshal(e); e == nil {
		println(string(bs))
	} else {
		t.Fatal(e)
	}
}
