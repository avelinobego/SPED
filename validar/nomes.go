package validar

import (
	"regexp"
	"strings"
)

const (
	CARACTERES_ESPECIAIS = `^([a-zA-Z]\s{0,1})*$`
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`^([a-zA-Z]\s{0,1})*[^\s]$`)
}

func ValidarNome(name string) bool {
	// Verifica se o nome tem mais de 3 caracteres e menos de 100 caracteres.
	if len(name) < 3 || len(name) > 100 {
		return false
	}

	partes := strings.Split(name, " ")
	if len(partes) < 2 {
		return false
	}

	// Verifica se o nome contém apenas letras e espaços.
	if !re.MatchString(name) {
		return false
	}

	return true
}
