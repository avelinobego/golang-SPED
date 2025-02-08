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

		result, error := xml.MarshalIndent(esocial, ".", "..")
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

	inclusao := builders.New[schemas.Inclusao]()

	infoEmpregador := builders.New[schemas.InfoEmpregador](
		builders.With("Inclusao", inclusao),
	)

	ideEmpregador := builders.New[schemas.IdeEmpregador](
		builders.With("NrInsc", "200"),
		builders.With("TpInsc", "1"),
	)

	evtInfoEmpregador := builders.New[schemas.EvtInfoEmpregador](
		builders.With("InfoEmpregador", infoEmpregador),
		builders.With("IdAttr", "1234567890"),
		builders.With("IdeEvento", "ID16944301890"),
		builders.With("IdeEmpregador", ideEmpregador),
	)

	var iface any = evtInfoEmpregador
	esocial = builders.New[schemas.ESocial](
		builders.With("Evento", &iface),
		builders.With("Namespace", "http://www.esocial.gov.br/schema/lote/eventos/envio/v1_2_3"),
	)

}
