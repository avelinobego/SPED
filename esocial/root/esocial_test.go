package root_test

import (
	"begoinformatica/sped/esocial/root"
	"encoding/xml"
	"fmt"
	"testing"
)

func TestRoot(t *testing.T) {
	root := &root.ESocial{}
	if out, err := xml.MarshalIndent(root, "\t", "\t"); err == nil {
		fmt.Println(string(out))
	} else {
		panic(err)
	}

}
