package validar

import (
	"begoinformatica/sped/comum"
	"fmt"
	"strings"
)

func make_pesos(start int) (result func() int) {
	peso := 2
	for i := 2; i <= start; i++ {
		if peso > 9 {
			peso = 2
		}
		peso++
	}

	result = func() (result int) {
		if peso == 1 {
			peso = 9

		}
		result = peso
		peso--
		return
	}
	return
}

func dig(n string) (resto int) {
	pesos := make_pesos(len(n))

	soma := 0
	for _, d := range n {
		soma = soma + (int(d-'0') * pesos())
	}

	resto = soma % 11

	if resto < 2 {
		resto = 0
	} else {
		resto = 11 - resto
	}
	return
}

func ValidarCNPJ(v string) (result bool) {

	result = false

	numero := comum.ApenasDigitos(v)
	if len(numero) < 14 {
		return
	}

	primeiro := string(numero[0])
	repetidos := strings.Repeat(primeiro, len(numero))
	if repetidos == numero {
		return
	}

	temp := numero[:len(numero)-2]
	d := dig(temp)
	temp = fmt.Sprintf("%s%d", temp, d)
	d = dig(temp)
	temp = fmt.Sprintf("%s%d", temp, d)

	result = numero == temp
	return
}
