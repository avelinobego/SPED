package util_test

import (
	"begoinformatica/sped/constantes"
	"begoinformatica/sped/util"
	"testing"
	"time"
)

func TestGerarId(t *testing.T) {
	t.Log("Comecando o teste")
	id := util.MakeId(constantes.CPF, 16944301890)
	for i := 0; i < 11; i++ {
		// time.Sleep(1 * time.Second)
		t.Log(id.Generate(time.Now()))
	}

}
