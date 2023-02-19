package validar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNomesValidos(t *testing.T) {
	assert.True(t, ValidarNome("Avelino Bego"), "falha")
}

func TestNomesInvalidos(t *testing.T) {
	assert.False(t, ValidarNome("Avelino@Bego"), "falha")
	assert.False(t, ValidarNome(" Avelino Bego"), "falha")
	assert.False(t, ValidarNome("Avelino@Bego"), "falha")
	assert.False(t, ValidarNome("Avelino Bego "), "falha")
	assert.False(t, ValidarNome("Avelino"), "falha")
}
