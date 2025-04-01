package utils_test

import (
	"encoding/xml"
	"log"
	"sped/utils"
	"testing"
)

type IdeEvento struct {
	Id_  int
	Nome string
}

type EvtInfoEmpregador struct {
	TpAmb_ int
	Ide    *IdeEvento
}

type ESocial struct {
	Ns_    string
	Id_    int
	Nome_  string
	Evento any
}

func TestCustomXml(t *testing.T) {

	idevt := &IdeEvento{
		Id_:  16944301890,
		Nome: "Avelino",
	}

	evt := &EvtInfoEmpregador{
		TpAmb_: 8,
		Ide:    idevt,
	}

	value := &ESocial{
		Ns_:    "http://www.esocial.gov.br/schema/lote/eventos/envio/v1",
		Id_:    1,
		Nome_:  "Avelino",
		Evento: evt,
	}

	decoded := utils.DecodeStruct(value)

	if v, e := xml.Marshal(decoded); e == nil {
		t.Logf("\n%s", v)
	} else {
		log.Fatal(e)
	}

}
