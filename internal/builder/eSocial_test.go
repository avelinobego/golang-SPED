package builder_test

import (
	"encoding/xml"
	"log"
	"os"
	"sped/internal/builder"
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

	inclusao, err := builder.New[schemas.Inclusao](
		builder.With("DadosRubrica", "dados"),
		builder.With("IdeRubrica", "1234"),
	)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	infoEmpregador, err := builder.New[schemas.InfoEmpregador](
		builder.With("Inclusao", inclusao))

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	ideEmpregador, err := builder.New[schemas.IdeEmpregador](
		builder.With("NrInsc", "200"),
		builder.With("TpInsc", "1"),
	)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	evtInfoEmpregador, err := builder.New[schemas.EvtInfoEmpregador](
		builder.With("InfoEmpregador", infoEmpregador),
		builder.With("IdAttr", "1234567890"),
		builder.With("IdeEvento", "ID16944301890"),
		builder.With("IdeEmpregador", ideEmpregador),
	)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	var iface any = evtInfoEmpregador
	esocial, err = builder.New[schemas.ESocial](
		builder.With("Evento", &iface),
		builder.With("Namespace", "http://www.esocial.gov.br/schema/lote/eventos/envio/v1_2_3"),
	)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

}
