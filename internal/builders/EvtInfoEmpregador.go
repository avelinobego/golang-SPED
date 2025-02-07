package builders

import (
	schemes "sped/internal/schemas"
)

type EvtInfoEmpregador schemes.EvtInfoEmpregador

type OptsEvtInfoEmpregador = func(o *EvtInfoEmpregador)

func NewEvtInfoEmpregador(opts ...OptsEvtInfoEmpregador) *EvtInfoEmpregador {

	o := new(EvtInfoEmpregador)

	for _, opt := range opts {
		opt(o)
	}

	return o
}

func WithIdAttr(value string) OptsEvtInfoEmpregador {
	return func(o *EvtInfoEmpregador) {
		o.IdAttr = value
	}
}

func (o *EvtInfoEmpregador) Validate() {
	o.IdAttr = ""
}
