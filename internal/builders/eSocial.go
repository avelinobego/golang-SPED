package builders

import (
	"fmt"
	"sped/internal/schemes"
)

// Makers para as classes dos schemes
func NewEsocial[T any](ns string, evento *T) (result *schemes.ESocial[T], err error) {
	if evento == nil {
		err = fmt.Errorf("evento nulo")
		return
	}
	result = &schemes.ESocial[T]{
		Namespace: ns,
		Evento:    evento,
	}
	return
}
