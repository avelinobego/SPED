package comum

import (
	"begoinformatica/sped/constantes"
	"encoding/xml"
)

type IdeEvento struct {
	XMLName xml.Name                         `xml:"ideEvento"`
	TpAmb   constantes.TipoAmbiente          `xml:"tpAmb"`
	ProcEmi constantes.ProcessoEmissaoEvento `xml:"procEmi"`
	VerProc string                           `xml:"verProc"`
}
