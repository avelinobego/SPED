package comum

import (
	"begoinformatica/sped/constantes"
	"encoding/xml"
)

type IdeEmpregador struct {
	XMLName xml.Name                 `xml:"ideEmpregador"`
	TpInsc  constantes.TipoInscricao `xml:"tpInsc"`
	NrInsc  string                   `xml:"nrInsc"`
}
