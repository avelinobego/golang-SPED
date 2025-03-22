package schema_test

import (
	"encoding/xml"
	"log"
	"os"
	"sped/internal/schema"
	"testing"

	"github.com/joho/godotenv"
)

var esocial interface{}

func TestMain(m *testing.M) {
	log.Println("Começo!")

	godotenv.Load("../../.env")

	exitVal := m.Run()
	if exitVal != 0 {
		log.Printf("Erro! %d", exitVal)
	} else {

		result, error := xml.MarshalIndent(esocial, " ", "  ")
		if error != nil {
			log.Fatal(error)
			return
		}

		log.Println(string(result))
		log.Println("Fim!")
	}
	os.Exit(exitVal)
}

func TestCreateEvtInfoEmpregador(t *testing.T) {

	inclusao := schema.Inclusao{
		DadosRubrica: "dados",
		IdeRubrica:   "1234",
	}

	infoEmpregador := schema.InfoEmpregador{
		Inclusao: &inclusao,
	}

	ideEmpregador := schema.IdeEmpregador{
		NrInsc: "200",
		TpInsc: "1",
	}

	evtInfoEmpregador := schema.EvtInfoEmpregador{
		InfoEmpregador: &infoEmpregador,
		IdAttr:         "1234567890",
		IdeEvento:      "ID16944301890",
		IdeEmpregador:  &ideEmpregador,
	}

	esocial = schema.ESocial[schema.EvtInfoEmpregador]{
		Namespace: "http://www.esocial.gov.br/schema/lote/eventos/envio/v1_2_3",
		Evento:    &evtInfoEmpregador,
	}

}
