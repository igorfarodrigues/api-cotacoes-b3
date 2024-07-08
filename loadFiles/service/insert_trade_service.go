package service

import (
	"log"
	"os"
	"path/filepath"

	"github.com/igorfarodrigues/api-cotacoes-b3/loadFiles/repository"
	"github.com/igorfarodrigues/api-cotacoes-b3/models"
	"github.com/igorfarodrigues/api-cotacoes-b3/utils"
)

func LoadAndSaveTrades(directory string) error {
	return filepath.Walk(directory, processFile)
}

func processFile(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if !info.IsDir() {
		if err := processTradesFromFile(path); err != nil {
			log.Printf("Erro ao processar arquivo %s: %v", path, err)
			return err
		}
	}
	return nil
}

func processTradesFromFile(filePath string) error {
	trades, err := utils.ReadTradesFromFile(filePath)
	if err != nil {
		return err
	}

	for _, trade := range trades {
		if err := saveTrade(trade); err != nil {
			return err
		}
	}

	return nil
}

func saveTrade(trade *models.Trade) error {
	if _, err := repository.SaveTrade(trade); err != nil {
		log.Printf("Erro ao salvar cotação: %v", err)
		return err
	}
	return nil
}
