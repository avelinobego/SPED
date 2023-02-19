package s1000

import (
	"begoinformatica/sped/constantes"
	"begoinformatica/sped/esocial/comum"
	"encoding/xml"
)

type EvtInfoEmpregador struct {
	XMLName            xml.Name `xml:"evtInfoEmpregador"`
	Id                 string   `xml:"id,attr"`
	IdeEvento          comum.IdeEvento
	IdePeriodo         IdePeriodo         `xml:"inclusao>idePeriodo"`
	InfoCadastro       InfoCadastro       `xml:"inclusao>infoCadastro"`
	IndAcordoIsenMulta constantes.TAcordo `xml:"infoOrgInternacional>indAcordoIsenMulta,omitempty"`
	Alteracao          Alteracao          `xml:"alteracao,omitempty"`
	Exclusao           IdePeriodo         `xml:"exclusao>idePeriodo,omitempty"`
}

type IdeEmpregador struct {
	XMLName xml.Name                 `xml:"ideEmpregador"`
	TpInsc  constantes.TipoInscricao `xml:"tpInsc"`
	NrInsc  string                   `xml:"nrInsc"`
}

type IdePeriodo struct {
	IniValid string `xml:"iniValid"`
	FimValid string `xml:"fimValid,omitempty"`
}

type InfoCadastro struct {
	XMLName               xml.Name                     `xml:"infoCadastro"`
	ClassTrib             string                       `xml:"classTrib"`
	IndCoop               constantes.TipoCooperativa   `xml:"indCoop,omitempty"`
	IndConstr             constantes.TipoConstrutora   `xml:"indConstr,omitempty"`
	IndDesFolha           constantes.TIndDesFolha      `xml:"indDesFolha"`
	IndOpcCP              constantes.TIndOpcCP         `xml:"indOpcCP,omitempty"`
	IndPorte              string                       `xml:"indPorte,omitempty"`
	IndOptRegEletron      constantes.TindOptRegEletron `xml:"indOptRegEletron"`
	CnpjEFR               string                       `xml:"cnpjEFR,omitempty"`
	DtTrans11096          comum.Data                   `xml:"dtTrans11096,omitempty"`
	IndTribFolhaPisCofins string                       `xml:"indTribFolhaPisCofins,omitempty"`
	DadosIsencao          DadosIsencao                 `xml:"dadosIsencao,omitempty"`
}

type DadosIsencao struct {
	IdeMinLei    string      `xml:"dadosIsencao"`
	NrCertif     string      `xml:"nrCertif"`
	DtEmisCertif *comum.Data `xml:"dtEmisCertif"`
	DtVencCertif *comum.Data `xml:"dtVencCertif"`
	NrProtRenov  string      `xml:"nrProtRenov,omitempty"`
	DtProtRenov  *comum.Data `xml:"dtProtRenov,omitempty"`
	DtDou        *comum.Data `xml:"dtDou,omitempty"`
	PagDou       int         `xml:"pagDou,omitempty"`
}

type Alteracao struct {
	IdePeriodo   IdePeriodo   `xml:"idePeriodo"`
	InfoCadastro InfoCadastro `xml:"infoCadastro"`
	NovaValidade IdePeriodo   `xml:"novaValidade,omitempty"`
}
