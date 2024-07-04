package repository

import (
	"database/sql"
	"github.com/igorfarodrigues/api-cotacoes-b3/models"
	_ "github.com/lib/pq"
)

type TradeRepository struct {
	DB *sql.DB
}

func NewTradeRepository(db *sql.DB) *TradeRepository {
	return &TradeRepository{DB: db}
}

func (repo *TradeRepository) SaveTrade(trade *models.Trade) error {
	query := `INSERT INTO trades (hora_fechamento, data_negocio, codigo_instrumento, preco_negocio, quantidade_negociada) VALUES ($1, $2, $3, $4, $5)`
	_, err := repo.DB.Exec(query, trade.HoraFechamento, trade.DataNegocio, trade.CodigoInstrumento, trade.PrecoNegocio, trade.QuantidadeNegociada)
	return err
}

func (repo *TradeRepository) GetMaxRangeValue(ticker string, date string) (float64, error) {
	var maxRangeValue float64
	query := `SELECT MAX(preco_negocio) FROM trades WHERE codigo_instrumento = $1 AND data_negocio >= $2`
	err := repo.DB.QueryRow(query, ticker, date).Scan(&maxRangeValue)
	return maxRangeValue, err
}

func (repo *TradeRepository) GetMaxDailyVolume(ticker string, date string) (int, error) {
	var maxDailyVolume int
	query := `SELECT MAX(total_volume) FROM (SELECT SUM(quantidade_negociada) as total_volume FROM trades WHERE codigo_instrumento = $1 AND data_negocio >= $2 GROUP BY data_negocio) as subquery`
	err := repo.DB.QueryRow(query, ticker, date).Scan(&maxDailyVolume)
	return maxDailyVolume, err
}
