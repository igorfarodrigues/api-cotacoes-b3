package repository

import (
	"fmt"
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


	// Verifique se o codigo_instrumento existe
	var exists bool
	err = conn.QueryRow("SELECT EXISTS (SELECT 1 FROM tickers WHERE codigo_instrumento = $1)", trade.CodigoInstrumento).Scan(&exists)
	if err != nil {
		return 0, fmt.Errorf("erro ao verificar existencia de codigo_instrumento: %v", err)
	}

	// Se não existir, insira na tabela de tickers
	if !exists {
		_, err := conn.Exec("INSERT INTO tickers (codigo_instrumento) VALUES ($1)", trade.CodigoInstrumento)
		if err != nil {
			return 0, fmt.Errorf("erro ao inserir novo codigo_instrumento: %v", err)
		}
	}



	query := `INSERT INTO trades (hora_fechamento, data_negocio, codigo_instrumento, preco_negocio, quantidade_negociada) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	if err = conn.QueryRow(query, trade.HoraFechamento, trade.DataNegocio, trade.CodigoInstrumento, trade.PrecoNegocio, trade.QuantidadeNegociada).Scan(&id); err != nil {
		log.Printf("Erro ao persistir os dados na tabela trades %d: %v", id, err)
	}
	return id, nil
}
