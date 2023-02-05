package comum_test

import (
	"begoinformatica/sped/comum"
	"fmt"
	"testing"
)

func TestApenasNumero(t *testing.T) {
	s := comum.ApenasDigitos("169.443.018-90")
	fmt.Println(s)
}
