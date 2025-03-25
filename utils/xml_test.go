package utils_test

import (
	"fmt"
	"sped/utils"
	"strings"
	"testing"
)

type IdeEvento struct {
	Id int
}

type EvtInfoEmpregador struct {
	TpAmb int
	Ide   *IdeEvento
}

type ESocial[T any] struct {
	Evento *T
}

func CleanBetween(initial string, start string, end string) string {
	result := strings.Builder{}
	include := true
	for _, v := range initial {
		if string(v) == start {
			include = false
		}
		if include {
			result.WriteRune(v)
		}
		if string(v) == end {
			include = true
		}
	}

	return result.String()
}

func TestCustomXml(t *testing.T) {

	evt := EvtInfoEmpregador{
		TpAmb: 8,
		Ide: &IdeEvento{
			Id: 16944301890,
		},
	}

	value := &ESocial[EvtInfoEmpregador]{
		Evento: &evt,
	}

	if v, e := utils.ToXml(value); e == nil {
		fmt.Println(v)
	}

}
