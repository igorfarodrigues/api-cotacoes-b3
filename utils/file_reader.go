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

	// Ignorar o cabeçalho
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var trades []*models.Trade
	lineNumber := 1 // Para contagem de linhas
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Erro ao ler o registro no arquivo %s na linha %d: %v", filepath, lineNumber, err)
			continue
		}

		log.Printf("Registro lido na linha %d: %v", lineNumber, record) // Adiciona log para o registro lido

		if len(record) < 10 {
			log.Printf("Registro incompleto no arquivo %s na linha %d: %v", filepath, lineNumber, record)
			lineNumber++ // Incrementa o número da linha
			continue     // Pula registros incompletos
		}

		precoNegocio, err := strconv.ParseFloat(strings.Replace(record[3], ",", ".", 1), 64)
		if err != nil {
			log.Printf("Erro ao parsear PrecoNegocio no arquivo %s na linha %d: %v", filepath, lineNumber, err)
			lineNumber++ // Incrementa o número da linha
			continue
		}

		qtdNegociada, err := strconv.Atoi(record[4])
		if err != nil {
			log.Printf("Erro ao parsear QuantidadeNegociada no arquivo %s na linha %d: %v", filepath, lineNumber, err)
			lineNumber++ // Incrementa o número da linha
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
		lineNumber++ // Incrementa o número da linha
	}
	return trades, nil
}
