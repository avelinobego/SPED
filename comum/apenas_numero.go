package comum

import (
	"strings"
	"unicode"
)

func ApenasDigitos(v string) (result string) {
	if v == "" {
		return
	}

	b := strings.Builder{}
	for _, char := range v {
		if !unicode.IsDigit(char) {
			continue
		}
		b.WriteByte(byte(char))
	}

	result = b.String()

	return
}
