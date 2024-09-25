package builders

import "sped/internal/schemes"

func WithId(id string) Option[schemes.EvtInfoEmpregador] {
	return func(value *schemes.EvtInfoEmpregador) {
		value.IdAttr = id
	}
}

func WithIdeEvento(id string) Option[schemes.EvtInfoEmpregador] {
	return func(value *schemes.EvtInfoEmpregador) {
		value.IdeEvento = id
	}
}

func WithIdeEmpregador(id *schemes.IdeEmpregador) Option[schemes.EvtInfoEmpregador] {
	return func(value *schemes.EvtInfoEmpregador) {
		value.IdeEmpregador = id
	}
}

func WithInfoEmpregador(id *schemes.InfoEmpregador) Option[schemes.EvtInfoEmpregador] {
	return func(value *schemes.EvtInfoEmpregador) {
		value.InfoEmpregador = id
	}
}

func NewEvtInfoEmpregador(opts ...Option[schemes.EvtInfoEmpregador]) (result *schemes.EvtInfoEmpregador) {
	result = &schemes.EvtInfoEmpregador{}
	for _, o := range opts {
		o(result)
	}
	return
}
