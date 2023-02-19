package constantes

// 1 - Contribuinte individual
// 2 - Produtor rural
// 3 - Segurado especial

type TipoCAEPF uint

const (
	Individual       TipoCAEPF = iota + 1
	ProdutorRural              = 2
	SeguradoEspecial           = 3
)
