package models

type Trade struct {
	HoraFechamento      string  `json:"hora_fechamento"`
	DataNegocio         string  `json:"data_negocio"`
	CodigoInstrumento   string  `json:"ticker"`
	PrecoNegocio        float64 `json:"preco_negocio"`
	QuantidadeNegociada int     `json:"qtd_negociada"`
}
