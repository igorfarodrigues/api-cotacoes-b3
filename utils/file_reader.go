package utils

import (
	"encoding/csv"
	"github.com/igorfarodrigues/api-cotacoes-b3/models"
	"os"
	"strconv"
	"strings"
)

func ReadTradesFromFile(filepath string) ([]*models.Trade, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var trades []*models.Trade
	for _, record := range records[1:] {
		precoNegocio, _ := strconv.ParseFloat(strings.Replace(record[3], ",", ".", -1), 64)
		quantidadeNegociada, _ := strconv.Atoi(record[4])
		trade := &models.Trade{
			CodigoInstrumento:   record[1],
			PrecoNegocio:        precoNegocio,
			QuantidadeNegociada: quantidadeNegociada,
			HoraFechamento:      record[5],
			DataNegocio:         record[8],
		}
		trades = append(trades, trade)
	}

	return trades, nil
}
