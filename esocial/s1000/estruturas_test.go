package s1000_test

import (
	"begoinformatica/sped/comum"
	"begoinformatica/sped/constantes"
	"begoinformatica/sped/esocial/root"
	"begoinformatica/sped/esocial/s1000"
	"encoding/xml"
	"fmt"
	"testing"
	"time"
)

func Check[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func TestMain(t *testing.T) {
	id := root.MakeId(constantes.CNPJ, 4126001000138)

	now := time.Now()

	info := s1000.EvtInfoEmpregador{Id: id.Generate(now)}
	ev := s1000.IdeEvento{}
	ip := s1000.IdePeriodo{}
	di := s1000.DadosIsencao{DtEmisCertif: comum.Data(now)}
	ic := s1000.InfoCadastro{DadosIsencao: di, IndCoop: 200}

	info.IdeEvento = ev
	info.IdePeriodo = ip
	info.InfoCadastro = ic

	r := root.ESocial{}
	r.Data = info

	if o, e := xml.MarshalIndent(r, "\t", "\t"); e == nil {
		fmt.Println(string(o))
	}

}
