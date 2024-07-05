package tests

import (
	"github.com/igorfarodrigues/api-cotacoes-b3/models"
	"github.com/igorfarodrigues/api-cotacoes-b3/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadTradesFromFile(t *testing.T) {

	content := `DataReferencia;CodigoInstrumento;AcaoAtualizacao;PrecoNegocio;QuantidadeNegociada;HoraFechamento;CodigoIdentificadorNegocio;TipoSessaoPregao;DataNegocio;CodigoParticipanteComprador;CodigoParticipanteVendedor
2024-07-03;WSPU24;0;5571,25;1;090001097;10;1;2024-07-03;8;120
2024-07-03;DI1F28;0;12,175;1;090000003;10;1;2024-07-03;120;114`

	tmpfile, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.WriteString(content)
	assert.NoError(t, err)

	err = tmpfile.Close()
	assert.NoError(t, err)

	trades, err := utils.ReadTradesFromFile(tmpfile.Name())
	assert.NoError(t, err)
	assert.Len(t, trades, 2)

	expectedTrades := []*models.Trade{
		{
			HoraFechamento:      "090001097",
			DataNegocio:         "2024-07-03",
			CodigoInstrumento:   "WSPU24",
			PrecoNegocio:        5571.25,
			QuantidadeNegociada: 1,
		},
		{
			HoraFechamento:      "090000003",
			DataNegocio:         "2024-07-03",
			CodigoInstrumento:   "DI1F28",
			PrecoNegocio:        12.175,
			QuantidadeNegociada: 1,
		},
	}

	for i, trade := range trades {
		assert.Equal(t, expectedTrades[i].HoraFechamento, trade.HoraFechamento)
		assert.Equal(t, expectedTrades[i].DataNegocio, trade.DataNegocio)
		assert.Equal(t, expectedTrades[i].CodigoInstrumento, trade.CodigoInstrumento)
		assert.Equal(t, expectedTrades[i].PrecoNegocio, trade.PrecoNegocio)
		assert.Equal(t, expectedTrades[i].QuantidadeNegociada, trade.QuantidadeNegociada)
	}
}
