package animal

import (
	enums "agrolumen/internal/domain/animal/enums"
	"time"
)

type Animal struct {
	// Individuais
	ID             int64      `db:"id" json:"id"`
	Brinco         *string    `db:"brinco" json:"brinco,omitempty"`
	Nome           *string    `db:"nome" json:"nome,omitempty"`
	Sexo           enums.Sexo `db:"sexo" json:"sexo"`
	DataNascimento time.Time  `db:"data_nascimento" json:"data_nascimento"`
	Raca           string     `db:"raca" json:"raca"`

	Origem            enums.Origem `db:"origem" json:"origem"`
	PropriedadeOrigem *string      `db:"propriedade_origem" json:"propriedade_origem,omitempty"`

	MaeID          *int64         `db:"mae_id" json:"mae_id,omitempty"`
	PaiID          *int64         `db:"pai_id" json:"pai_id,omitempty"`
	GrupoSanguineo *string        `db:"grupo_sanguineo" json:"grupo_sanguineo,omitempty"`
	Situacao       enums.Situacao `db:"situacao" json:"situacao"`
	DataSaida      *time.Time     `db:"data_saida" json:"data_saida,omitempty"`
	MotivoSaida    *string        `db:"motivo_saida" json:"motivo_saida,omitempty"`

	//Produção (válido para fêmeas)
	PrimeiraLactacao    *time.Time `db:"primeira_lactacao" json:"primeira_lactacao,omitempty"`
	ProducaoDia         *float64   `db:"producao_dia" json:"producao_dia,omitempty"`
	ProducaoMediaMensal *float64   `db:"producao_media_mensal" json:"producao_medial_mensal,omitempty"`
	TotalLactacoes      *int64     `db:"total_lactacoes" json:"total_lactacoes,omitempty"`
	PicoProducao        *float64   `db:"pico_producao" json:"pico_producao,omitempty"`
	DiasEmLactacao      *int64     `db:"dias_em_lactacao" json:"dias_em_lactacao,omitempty"`

	//Reprodução
	Reproducao      *enums.Reproducao `db:"reproducao" json:"reproducao,omitempty"`
	UltimaCobertura *time.Time        `db:"ultima_cobertura" json:"ultima_cobertura,omitempty"`
	Cobertura       *enums.Cobertura  `db:"cobertura" json:"cobertura,omitempty"`
	Reprodutor      *string           `db:"reprodutor" json:"reprodutor,omitempty"`
	Gestacao        *enums.Gestacao   `db:"gestacao" json:"gestacao,omitempty"`
	DataDiagnostico *time.Time        `db:"data_diagnostico" json:"data_diagnostico,omitempty"`
	PrevisaoParto   *time.Time        `db:"previsao_parto" json:"previsao_parto,omitempty"`
	TotalPartos     *int64            `db:"total_partos" json:"total_partos,omitempty"`
	UltimoParto     *time.Time        `db:"ultimo_parto" json:"ultimo_parto,omitempty"`
	IntervaloPartos *int64            `db:"intervalo_partos" json:"intervalo_partos,omitempty"`

	//Sanitário
	Sanitario        enums.Sanitario `db:"sanitario" json:"sanitario"`
	UltimoVermifugo  *time.Time      `db:"ultima_vermifugacao" json:"ultima_vermifugacao,omitempty"`
	UltimaVacina     *time.Time      `db:"utlima_vacina" json:"ultima_vacina,omitempty"`
	HistoricoVacinas *string         `db:"historico_vacinas" json:"historico_vacinas,omitempty"`
	HistoricoDoencas *string         `db:"historico_doencas" json:"historico_doencas,omitempty"`
	ExamesRealizados *string         `db:"exames_realizados" json:"exames_realizados,omitempty"`

	//Peso e Escore
	PesoAtual   *float64   `db:"peso_atual" json:"peso_atual,omitempty"`
	Escore      *float64   `db:"escore" json:"escore,omitempty"`
	DataPesagem *time.Time `db:"data_pesagem" json:"data_pesagem,omitempty"`

	//Outros
	Observacoes *string `db:"observacoes" json:"observacoes,omitempty"`
	FotoURL     *string `db:"foto_url" json:"foto_url,omitempty"`

	//Auditoria
	UserID    int64      `db:"user_id" json:"user_id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
