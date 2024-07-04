package tests

import (
	"github.com/igorfarodrigues/api-cotacoes-b3/models"
	"github.com/igorfarodrigues/api-cotacoes-b3/repository"
	"github.com/igorfarodrigues/api-cotacoes-b3/service"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTradeData(t *testing.T) {
	// Adicionar dados de teste
	trade := &models.Trade{
		CodigoInstrumento:   "WSPU24",
		PrecoNegocio:        5571.25,
		QuantidadeNegociada: 1,
		HoraFechamento:      "090001097",
		DataNegocio:         "2024-07-03",
	}
	_,_ = repository.SaveTrade(trade)

	// Testar função
	data, err := service.GetTradeData("WSPU24", "2024-07-03")
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, 5571.25, data["max_range_value"])
}
