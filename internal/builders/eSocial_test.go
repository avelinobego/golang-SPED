package builders_test

import (
	"encoding/xml"
	"log"
	"os"
	"sped/internal/builders"
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

	inclusao, err := builders.New[schemas.Inclusao](
		builders.With("DadosRubrica", "dados"),
		builders.With("IdeRubrica", "1234"),
	)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	infoEmpregador, err := builders.New[schemas.InfoEmpregador](
		builders.With("Inclusao", inclusao))

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	ideEmpregador, err := builders.New[schemas.IdeEmpregador](
		builders.With("NrInsc", "200"),
		builders.With("TpInsc", "1"),
	)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	evtInfoEmpregador, err := builders.New[schemas.EvtInfoEmpregador](
		builders.With("InfoEmpregador", infoEmpregador),
		builders.With("IdAttr", "1234567890"),
		builders.With("IdeEvento", "ID16944301890"),
		builders.With("IdeEmpregador", ideEmpregador),
	)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	var iface any = evtInfoEmpregador
	esocial, err = builders.New[schemas.ESocial](
		builders.With("Evento", &iface),
		builders.With("Namespace", "http://www.esocial.gov.br/schema/lote/eventos/envio/v1_2_3"),
	)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

}
