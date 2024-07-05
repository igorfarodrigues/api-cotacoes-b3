package utils

import (
	"encoding/csv"
	"github.com/igorfarodrigues/api-cotacoes-b3/models"
	"io"
	"log"
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
	reader.TrimLeadingSpace = true

	var trades []*models.Trade
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(record) < 10 {
			continue // Pula registros incompletos
		}

		precoNegocio, err := strconv.ParseFloat(strings.Replace(record[3], ",", ".", 1), 64)
		if err != nil {
			log.Printf("Erro ao parsear PrecoNegocio no arquivo %s: %v", filepath, err)
			continue
		}

		qtdNegociada, err := strconv.Atoi(record[4])
		if err != nil {
			log.Printf("Erro ao parsear QuantidadeNegociada no arquivo %s: %v", filepath, err)
			continue
		}

		trade := &models.Trade{
			HoraFechamento:      record[5],
			DataNegocio:         record[8],
			CodigoInstrumento:   record[1],
			PrecoNegocio:        precoNegocio,
			QuantidadeNegociada: qtdNegociada,
		}
		trades = append(trades, trade)
	}
	return trades, nil
}
