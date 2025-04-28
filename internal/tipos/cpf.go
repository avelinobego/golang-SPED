package tipos

import (
	"fmt"
	"sped/utils"
)

type Cpf string

func (c Cpf) MarshalText() (text []byte, err error) {
	if !utils.ValidaCPF(string(c)) {
		return nil, fmt.Errorf("valor de CPF inválido: %s", c)
	}
	return []byte(c), nil
}
