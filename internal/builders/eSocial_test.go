package builders_test

import (
	"encoding/xml"
	"log"
	"os"
	"sped/internal/builders"
	"sped/internal/schemes"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Começo!")
	exitVal := m.Run()
	if exitVal != 0 {
		log.Printf("Erro! %d", exitVal)
	} else {
		log.Println("Fim!")
	}
	os.Exit(exitVal)
}

func TestESocial(t *testing.T) {

	ide_empregador := &schemes.IdeEmpregador{}
	info_empregador := &schemes.InfoEmpregador{}

	evt := builders.NewEvtInfoEmpregador(
		builders.WithInfoEmpregador(info_empregador),
		builders.WithIdeEmpregador(ide_empregador),
		builders.WithIdeEvento("ide_evento_123"),
		builders.WithId("1234567890"))

	esocial, err := builders.NewEsocial("xpto", evt)
	if err != nil {
		panic(err)
	}

	text, err := xml.MarshalIndent(esocial, " ", "   ")
	if err != nil {
		panic(err)
	}

	t.Logf("\n%s", string(text))
}
