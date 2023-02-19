package s1000

import (
	"begoinformatica/sped/esocial/comum"
)

func Make[T any]() (esocial *comum.ESocial[T]) {
	esocial = &comum.ESocial[T]{}
	return
}
