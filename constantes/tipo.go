package constantes

type Tipo int64

const (
	CPF Tipo = iota + 1
	CNPJ
)

type TipoPessoa uint

const (
	Fisica TipoPessoa = iota + 1
	Juridica
)

// Identificação do ambiente.
// Valores válidos:
// 1 - Produção
// 2 - Produção restrita
// 7 - Validação (uso interno)
// 8 - Teste (uso interno)
// 9 - Desenvolvimento (uso interno)

type TipoAmbiente uint

const (
	Producao         TipoAmbiente = iota + 1
	ProducaoRestrita              = 2
	Validacao                     = 7
	Teste                         = 8
	Desenvolvimento               = 9
)

// Processo de emissão do evento.
// Valores válidos:
// 1 - Aplicativo do empregador
// 2 - Aplicativo governamental - Simplificado Pessoa Física
// 3 - Aplicativo governamental - Web Geral
// 4 - Aplicativo governamental - Simplificado Pessoa Jurídica
// 9 - Aplicativo governamental - Integração com a Junta Comercial
// 22 - Aplicativo governamental para dispositivos móveis - Empregador Doméstico

type ProcessoEmissaoEvento uint

const (
	Empregador     ProcessoEmissaoEvento = iota + 1
	SimplificadoPF                       = 2
	WebGeral                             = 3
	SimplificadoPJ                       = 4
	IntegracaoJC                         = 9
	Domestico                            = 22
)

type TipoInscricao uint

const (
	InscrCNPJ  TipoInscricao = iota + 1
	InscrCPF                 = 2
	InscrCAEPF               = 3
	InscrCNO                 = 4
)

// 0 - Não é cooperativa
// 1 - Cooperativa de trabalho
// 2 - Cooperativa de produção
// 3 - Outras cooperativas

type TipoCooperativa uint

const (
	CoopNao      TipoCooperativa = iota
	CoopTrabalho                 = 1
	CoopProducao                 = 2
	CoopOutras                   = 3
)

type TipoConstrutora uint

const (
	NaoConstrutora TipoConstrutora = iota
	Construtora                    = 1
)

type TIndDesFolha uint

const (
	NaoAplicavel      TIndDesFolha = iota
	EmpresaEnquadrada              = 1
)
