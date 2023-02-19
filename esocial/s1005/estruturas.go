package s1005

import (
	"begoinformatica/sped/constantes"
	"begoinformatica/sped/esocial/comum"
	"encoding/xml"
)

type EvtTabEstab struct {
	XMLName    xml.Name `xml:"evtTabEstab"`
	Id         string   `xml:"id,attr"`
	IdeEvent   comum.IdeEvento
	IdeEstab   IdeEstab   `xml:"inclusao>evtTabEstab,omitempty"`
	DadosEstab DadosEstab `xml:"inclusao>dadosEstab"`
}

type IdeEstab struct {
	XMLName  xml.Name                 `xml:"evtTabEstab"`
	TpInsc   constantes.TipoInscricao `xml:"tpInsc"`
	NrInsc   string                   `xml:"nrInsc"`
	IniValid comum.Data               `xml:"iniValid"`
	FimValid comum.Data               `xml:"fimValid,omitempty"`
}

type DadosEstab struct {
	XMLName    xml.Name   `xml:"dadosEstab"`
	CnaePrep   int        `xml:"cnaePrep"`
	CnpjResp   string     `xml:"cnpjResp"`
	AliqGilrat AliqGilrat `xml:",omitempty"`
	InfoCaepf  InfoCaepf  `xml:",omitempty"`
	InfoObra   InfoObra   `xml:",omitempty"`
}

type AliqGilrat struct {
	XMLName       xml.Name               `xml:"aliqGilrat"`
	AliqRat       constantes.AliquotaRAT `xml:"aliqRat"`
	Fap           int                    `xml:"fap"`
	ProcAdmJudRat ProcAdmJudRat          `xml:",omitempty"`
	ProcAdmJudFap ProcAdmJudFap          `xml:",omitempty"`
}

type ProcAdmJudRat struct {
	XMLName xml.Name                  `xml:"procAdmJudRat"`
	TpProc  constantes.TiposProcessos `xml:"tpProc"`
	NrProc  string                    `xml:"nrProc"`
	CodSusp int                       `xml:"codSusp"`
}

type ProcAdmJudFap struct {
	XMLName xml.Name                  `xml:"procAdmJudFap"`
	TpProc  constantes.TiposProcessos `xml:"tpProc"`
	NrProc  string                    `xml:"nrProc"`
	CodSusp int                       `xml:"codSusp"`
}

type InfoCaepf struct {
	XMLName xml.Name             `xml:"infoCaepf"`
	TpCaepf constantes.TipoCAEPF `xml:"tpCaepf"`
}

type InfoObra struct {
	XMLName          xml.Name                    `xml:"indSubstPatrObra"`
	IndSubstPatrObra constantes.TContribPatronal `xml:"indSubstPatrObra"`
}
