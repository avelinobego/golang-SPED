package util_test

import (
	"testing"

	"github.com/avelinobego/sped/util"
)

func TestInverter(t *testing.T) {
	numero := "16944301890"
	for i := len(numero) - 1; i > -1; i-- {
		t.Log(string(numero[i]))
	}
}

func TestPesos(t *testing.T) {
	cnpj := "04.126.001/0001-38"
	// cnpj := "11.111.111/1111-12"

	if util.ValidarCNPJ(cnpj) {
		t.Log("correto")
	} else {
		t.Log("invalido")
	}

}
