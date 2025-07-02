package animal

import (
	"agrolumen/internal/domain/animal/enums"
	"time"
)

type AnimalBaseDTO struct {
	ID             int64          `json:"id"`
	Nome           *string        `json:"nome,omitempty"`
	Brinco         *string        `json:"brinco,omitempty"`
	Sexo           enums.Sexo     `json:"sexo"`
	DataNascimento time.Time      `json:"data_nascimento"`
	Raca           string         `json:"raca"`
	GrupoSanguineo *string        `json:"grupo_sanguineo,omitempty"`
	Situacao       enums.Situacao `json:"situacao"`
	DataSaida      *time.Time     `json:"data_saida,omitempty"`
	MotivoSaida    *string        `json:"motivo_saida,omitempty"`

	Observacoes *string `json:"observacoes,omitempty"`
	FotoURL     *string `json:"foto_url,omitempty"`
}
type AnimalOrigemDTO struct {
	Origem            enums.Origem `json:"origem"`
	PropriedadeOrigem *string      `json:"propriedade_origem,omitempty"`
}
type AnimalGenealogiaDTO struct {
	MaeID *int64 `json:"mae_id,omitempty"`
	PaiID *int64 `json:"pai_id,omitempty"`
}

type AnimalProducaoDTO struct {
	PrimeiraLactacao    *time.Time `json:"primeira_lactacao,omitempty"`
	ProducaoDia         *float64   `json:"producao_dia,omitempty"`
	ProducaoMediaMensal *float64   `json:"producao_medial_mensal,omitempty"`
	TotalLactacoes      *int64     `json:"total_lactacoes,omitempty"`
	PicoProducao        *float64   `json:"pico_producao,omitempty"`
	DiasEmLactacao      *int64     `json:"dias_em_lactacao,omitempty"`
}

type AnimalReproducaoDTO struct {
	Reproducao      *enums.Reproducao `json:"reproducao,omitempty"`
	UltimaCobertura *time.Time        `json:"ultima_cobertura,omitempty"`
	Cobertura       *enums.Cobertura  `json:"cobertura,omitempty"`
	Reprodutor      *string           `json:"reprodutor,omitempty"`
	Gestacao        *enums.Gestacao   `json:"gestacao,omitempty"`
	DataDiagnostico *time.Time        `json:"data_diagnostico,omitempty"`
	PrevisaoParto   *time.Time        `json:"previsao_parto,omitempty"`
	TotalPartos     *int64            `json:"total_partos,omitempty"`
	UltimoParto     *time.Time        `json:"ultimo_parto,omitempty"`
	IntervaloPartos *int64            `json:"intervalo_partos,omitempty"`
}

type AnimalSanitarioDTO struct {
	Sanitario        enums.Sanitario `json:"sanitario"`
	UltimoVermifugo  *time.Time      `json:"ultima_vermifugacao,omitempty"`
	UltimaVacina     *time.Time      `json:"ultima_vacina,omitempty"`
	HistoricoVacinas *string         `json:"historico_vacinas,omitempty"`
	HistoricoDoencas *string         `json:"historico_doencas,omitempty"`
	ExamesRealizados *string         `json:"exames_realizados,omitempty"`
}

type AnimalPesoDTO struct {
	PesoAtual   *float64   `json:"peso_atual,omitempty"`
	Escore      *float64   `json:"escore,omitempty"`
	DataPesagem *time.Time `json:"data_pesagem,omitempty"`
}
