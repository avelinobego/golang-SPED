package schemas_test

import (
	"encoding/xml"
	"log"
	"os"
	"sped/internal/schemas"
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

	inclusao := schemas.Inclusao{
		DadosRubrica: "dados",
		IdeRubrica:   "1234",
	}

	infoEmpregador := schemas.InfoEmpregador{
		Inclusao: &inclusao,
	}

	ideEmpregador := schemas.IdeEmpregador{
		NrInsc: "200",
		TpInsc: "1",
	}

	evtInfoEmpregador := schemas.EvtInfoEmpregador{
		InfoEmpregador: &infoEmpregador,
		IdAttr:         "1234567890",
		IdeEvento:      "ID16944301890",
		IdeEmpregador:  &ideEmpregador,
	}

	esocial = schemas.ESocial[schemas.EvtInfoEmpregador]{
		Namespace: "http://www.esocial.gov.br/schema/lote/eventos/envio/v1_2_3",
		Evento:    &evtInfoEmpregador,
	}

}
