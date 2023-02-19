package util

import (
	"fmt"
	"testing"
)

func TestApenasNumero(t *testing.T) {
	s := ApenasDigitos("169.443.018-90")
	fmt.Println(s)
}
