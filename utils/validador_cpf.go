package utils

import (
	"regexp"
	"strconv"
)

// ValidaCPF valida um CPF em formato string
func ValidaCPF(cpf string) bool {
	// Remover tudo que não é número
	re := regexp.MustCompile(`[^0-9]`)
	cpf = re.ReplaceAllString(cpf, "")

	if len(cpf) != 11 {
		return false
	}

	// CPF com todos dígitos iguais é inválido
	todosIguais := true
	for i := 1; i < 11 && todosIguais; i++ {
		if cpf[i] != cpf[0] {
			todosIguais = false
		}
	}
	if todosIguais {
		return false
	}

	// Calcular o primeiro dígito verificador
	soma := 0
	for i := 0; i < 9; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		soma += num * (10 - i)
	}
	digito1 := (soma * 10) % 11
	if digito1 == 10 {
		digito1 = 0
	}

	// Verifica o primeiro dígito
	if digito1 != int(cpf[9]-'0') {
		return false
	}

	// Calcular o segundo dígito verificador
	soma = 0
	for i := range 10 {
		num, _ := strconv.Atoi(string(cpf[i]))
		soma += num * (11 - i)
	}
	digito2 := (soma * 10) % 11
	if digito2 == 10 {
		digito2 = 0
	}

	// Verifica o segundo dígito
	return digito2 == int(cpf[10]-'0')
}
