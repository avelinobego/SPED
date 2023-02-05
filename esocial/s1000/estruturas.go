package s1000

import (
	"begoinformatica/sped/comum"
	"begoinformatica/sped/constantes"
	"encoding/xml"
)

type EvtInfoEmpregador struct {
	XMLName            xml.Name `xml:"evtInfoEmpregador"`
	Id                 string   `xml:"id,attr"`
	IdeEvento          IdeEvento
	IdePeriodo         IdePeriodo   `xml:"inclusao>idePeriodo"`
	InfoCadastro       InfoCadastro `xml:"inclusao>infoCadastro"`
	IndAcordoIsenMulta comum.Numero `xml:"infoOrgInternacional>indAcordoIsenMulta,omitempty"`
	Alteracao          Alteracao    `xml:"alteracao,omitempty"`
	Exclusao           IdePeriodo   `xml:"exclusao>idePeriodo,omitempty"`
}

type IdeEvento struct {
	XMLName xml.Name                         `xml:"ideEvento"`
	TpAmb   constantes.TipoAmbiente          `xml:"tpAmb"`
	ProcEmi constantes.ProcessoEmissaoEvento `xml:"procEmi"`
	VerProc string                           `xml:"verProc"`
}

type IdeEmpregador struct {
	XMLName xml.Name     `xml:"ideEmpregador"`
	TpInsc  comum.Numero `xml:"tpInsc"`
	NrInsc  string       `xml:"nrInsc"`
}

type IdePeriodo struct {
	IniValid string `xml:"iniValid"`
	FimValid string `xml:"fimValid,omitempty"`
}

type InfoCadastro struct {
	XMLName               xml.Name     `xml:"infoCadastro"`
	ClassTrib             string       `xml:"classTrib"`
	IndCoop               comum.Numero `xml:"indCoop,omitempty"`
	IndConstr             comum.Numero `xml:"indConstr,omitempty"`
	IndDesFolha           comum.Numero `xml:"indDesFolha"`
	IndOpcCP              comum.Numero `xml:"indOpcCP,omitempty"`
	IndPorte              string       `xml:"indPorte,omitempty"`
	IndOptRegEletron      comum.Numero `xml:"indOptRegEletron"`
	CnpjEFR               string       `xml:"cnpjEFR,omitempty"`
	DtTrans11096          comum.Data   `xml:"dtTrans11096,omitempty"`
	IndTribFolhaPisCofins string       `xml:"indTribFolhaPisCofins,omitempty"`
	DadosIsencao          DadosIsencao `xml:"dadosIsencao,omitempty"`
}

type DadosIsencao struct {
	IdeMinLei    string       `xml:"dadosIsencao"`
	NrCertif     string       `xml:"nrCertif"`
	DtEmisCertif comum.Data   `xml:"dtEmisCertif"`
	DtVencCertif comum.Data   `xml:"dtVencCertif"`
	NrProtRenov  string       `xml:"nrProtRenov,omitempty"`
	DtProtRenov  comum.Data   `xml:"dtProtRenov,omitempty"`
	DtDou        comum.Data   `xml:"dtDou,omitempty"`
	PagDou       comum.Numero `xml:"pagDou,omitempty"`
}

type Alteracao struct {
	IdePeriodo   IdePeriodo   `xml:"idePeriodo"`
	InfoCadastro InfoCadastro `xml:"infoCadastro"`
	NovaValidade IdePeriodo   `xml:"novaValidade,omitempty"`
}
