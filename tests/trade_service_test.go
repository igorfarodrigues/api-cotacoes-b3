package tests

import (
	"database/sql"
	"github.com/igorfarodrigues/api-cotacoes-b3/models"
	"github.com/igorfarodrigues/api-cotacoes-b3/repository"
	"github.com/igorfarodrigues/api-cotacoes-b3/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTradeData(t *testing.T) {
	db, _ := sql.Open("postgres", "user=youruser password=yourpassword dbname=yourdb sslmode=disable")
	repo := repository.NewTradeRepository(db)
	tradeService := service.NewTradeService(repo)

	// Adicionar dados de teste
	trade := &models.Trade{
		HoraFechamento:      "17:00:00",
		DataNegocio:         "2023-12-01",
		CodigoInstrumento:   "PETR4",
		PrecoNegocio:        20.0,
		QuantidadeNegociada: 10,
	}
	tradeService.SaveTrade(trade)

	data, err := tradeService.GetTradeData("PETR4", "2023-12-01")
	assert.NoError(t, err)
	assert.Equal(t, "PETR4", data["ticker"])
	assert.Equal(t, 20.0, data["max_range_value"])
	assert.Equal(t, 10, data["max_daily_volume"])
}
