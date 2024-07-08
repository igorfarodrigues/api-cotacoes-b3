package repository

import (
	"log"

	"github.com/igorfarodrigues/api-cotacoes-b3/db"
	"github.com/igorfarodrigues/api-cotacoes-b3/models"
	_ "github.com/lib/pq"
)

func SaveTrade(trade *models.Trade) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	query := `INSERT INTO trades (hora_fechamento, data_negocio, codigo_instrumento, preco_negocio, quantidade_negociada) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	if err = conn.QueryRow(query, trade.HoraFechamento, trade.DataNegocio, trade.CodigoInstrumento, trade.PrecoNegocio, trade.QuantidadeNegociada).Scan(&id); err != nil {
		log.Printf("Erro ao persistir os dados na tabela trades %d: %v", id, err)
	}
	return
}
