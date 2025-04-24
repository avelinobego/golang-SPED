package tipos_test

import (
	"encoding/xml"
	"log"
	"sped/internal/tipos"
	"testing"
	"time"
)

type Container struct {
	Id     tipos.Id   `xml:"id,attr"`
	Comeco tipos.Data `xml:"dtIni"`
	Fim    tipos.Data `xml:"dtFim"`
}

func TestFormat(t *testing.T) {
	comeco := tipos.Data(time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC))
	fim := tipos.Data(time.Date(2030, time.December, 1, 0, 0, 0, 0, time.UTC))

	c := Container{
		Id:     tipos.Id{},
		Comeco: comeco,
		Fim:    fim,
	}

	if bs, e := xml.Marshal(c); e != nil {
		log.Fatal(e)
	} else {
		log.Println(string(bs))
	}

}
