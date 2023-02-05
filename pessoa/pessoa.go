package pessoa

import (
	"begoinformatica/sped/constantes"
	"time"
)

type Pessoa struct {
	Tipo           constantes.TipoPessoa
	Nome           string
	DataNascimento *time.Time
	CPF            int32
	CNPJ           int32
	Nascimento     *time.Time
	Telefone       string
	Email          string
}
