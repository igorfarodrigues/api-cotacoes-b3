package repository

import (
	"github.com/igorfarodrigues/api-cotacoes-b3/db"
	_ "github.com/lib/pq"
)

func GetMaxRangeValue(ticker string, date string) (maxRangeValue float64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return maxRangeValue, err
	}

	defer conn.Close()

	query := `SELECT MAX(preco_negocio) FROM trades WHERE codigo_instrumento = $1 AND data_negocio >= $2`
	err = conn.QueryRow(query, ticker, date).Scan(&maxRangeValue)
	return maxRangeValue, err
}

func GetMaxDailyVolume(ticker string, date string) (maxDailyVolume int, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return maxDailyVolume, err
	}

	defer conn.Close()

	query := `SELECT MAX(total_volume) FROM (SELECT SUM(quantidade_negociada) as total_volume FROM trades WHERE codigo_instrumento = $1 AND data_negocio >= $2 GROUP BY data_negocio) as subquery`
	err = conn.QueryRow(query, ticker, date).Scan(&maxDailyVolume)
	return maxDailyVolume, err
}
