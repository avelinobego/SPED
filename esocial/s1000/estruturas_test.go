package s1000_test

import (
	"begoinformatica/sped/constantes"
	"begoinformatica/sped/esocial/comum"
	"begoinformatica/sped/esocial/s1000"
	"begoinformatica/sped/util"
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
	id := util.MakeId(constantes.CNPJ, 4126001000138)
	now := time.Now()

	info := s1000.EvtInfoEmpregador{Id: id.Generate(now)}
	ev := comum.IdeEvento{}
	ip := s1000.IdePeriodo{}
	di := s1000.DadosIsencao{DtEmisCertif: comum.NewData(now)}
	ic := s1000.InfoCadastro{DadosIsencao: di, IndCoop: 200}

	info.IdeEvento = ev
	info.IdePeriodo = ip
	info.InfoCadastro = ic

	r := comum.ESocial[s1000.EvtInfoEmpregador]{}
	r.Data = info

	if o, e := xml.MarshalIndent(r, "\t", "\t"); e == nil {
		fmt.Println(string(o))
	}

}
