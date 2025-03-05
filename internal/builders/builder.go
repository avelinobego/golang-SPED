package builders

import (
	"fmt"
	"reflect"
)

type Opts func(s any) error

func New[evento any](opts ...Opts) (evt *evento, e error) {

	tipoGenerico := reflect.TypeOf((*evento)(nil)).Elem().Kind()
	if tipoGenerico != reflect.Struct {
		e = fmt.Errorf("o parâmetro genérico precisa ser do tipo struct e não %s", tipoGenerico)
		return
	}

	var result evento
	for _, f := range opts {
		e = f(&result)
		if e != nil {
			return
		}
	}
	evt = &result
	return
}

func With(name string, value any) Opts {
	return func(s any) (e error) {
		t := reflect.ValueOf(s).Elem()
		f := t.FieldByName(name)

		if !f.IsValid() {
			e = fmt.Errorf("propriedade inválida: %s", name)
		} else if !f.CanSet() {
			e = fmt.Errorf("não posso atribir à propriedade: %s", name)
		} else {
			inputValue := reflect.ValueOf(value)
			if f.Type() == inputValue.Type() {
				f.Set(inputValue)
			} else {
				e = fmt.Errorf("tipos incompatíveis: %s esperava %s mas recebeu %s em %s", name, f.Type(), inputValue.Type(), t.Type().Name())
			}
		}

		return
	}

}
