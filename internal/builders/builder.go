package builders

import (
	"reflect"
)

type Opts func(s any)

func New[evento any](opts ...Opts) *evento {

	var result evento

	for _, f := range opts {
		f(&result)
	}

	return &result
}

func With(name string, value any) Opts {
	return func(s any) {
		t := reflect.ValueOf(s).Elem()
		f := t.FieldByName(name)

		if !f.IsValid() {
			panic("propriedade inválida")
		}

		if !f.CanSet() {
			panic("não posso atribir à propriedade")
		}

		v := reflect.ValueOf(value)
		if f.Kind() == v.Kind() {
			f.Set(v)
		} else {
			panic("tipos incompatíveis")
		}

	}
}
