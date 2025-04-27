package events

import "sped/internal/tipos"

type EvtAdmissao struct {
	XMLName       struct{}                `xml:"evtAdmissao"`
	IdeEvento     tipos.IdeEventoAdmissao `xml:"ideEvento"`
	IdeEmpregador tipos.IdeEmpregador     `xml:"ideEmpregador"`
	Trabalhador   tipos.Trabalhador       `xml:"trabalhador"`
}
