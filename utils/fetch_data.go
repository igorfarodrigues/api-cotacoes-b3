package utils

import (
	"encoding/csv"
	"github.com/igorfarodrigues/api-cotacoes-b3/models"
	"net/http"
	"strconv"
	"strings"
)

func FetchData(url string) ([]*models.Trade, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
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
			HoraFechamento:      record[0],
			DataNegocio:         record[1],
			CodigoInstrumento:   record[2],
			PrecoNegocio:        precoNegocio,
			QuantidadeNegociada: quantidadeNegociada,
		}
		trades = append(trades, trade)
	}

	return trades, nil
}
