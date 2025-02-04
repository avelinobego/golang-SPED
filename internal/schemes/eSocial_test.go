package schemes_test

import (
	"encoding/xml"
	"log"
	"os"
	"sped/internal/schemes"
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

	evtInfo := schemes.NewEvtInfoEmpregador(
		schemes.WithInfoEmpIdAttr("1234567890"),
		schemes.WithInfoEmpIdeEmpregador(&schemes.IdeEmpregador{}),
		schemes.WithInfoEmpIdeEvento("1234567890"),
		schemes.WithInfoEmpInfoEmpregador(&schemes.InfoEmpregador{}),
	)

	// content, error := xml.MarshalIndent(&evtInfo, ".", "..")
	// if error != nil {
	// 	log.Fatal(error)
	// }
	// log.Println(string(content))

	esocial := schemes.NewESocial[schemes.EvtInfoEmpregador](
		schemes.WithESocialNamespace[schemes.EvtInfoEmpregador]("xpto"),
		schemes.WithESocialEvento[schemes.EvtInfoEmpregador](&evtInfo),
	)

	// error = esocial.Validate()
	// if error != nil {
	// 	log.Fatal(error)
	// }

	content, error := xml.MarshalIndent(&esocial, ".", "..")
	if error != nil {
		log.Fatal(error)
	}
	log.Println(string(content))

}
