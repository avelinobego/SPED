package data

type CategTrabalhadores struct {
	Id         string
	Descricao  string
	Grupo      string
	CP         string
	Codigo     string
	DtInicio   string
	DtFim      string
	Aliqfgts   string
	Obrigacao  string
	Aliqfgtsco string
}

func (c *CategTrabalhadores) Load(id string) {

}
