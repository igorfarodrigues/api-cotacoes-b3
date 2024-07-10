package main

import (
	"log"

	"github.com/igorfarodrigues/api-cotacoes-b3/configs"
	"github.com/igorfarodrigues/api-cotacoes-b3/pkg/service"
)

var jobs = make(chan string, 100)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Log the loaded configuration
	log.Printf("Loaded configuration: %+v", configs.GetFolderPath())
	jobs <- configs.GetFolderPath()
	go worker() // Inicia o worker para processar os arquivos

	err = worker()
	if err != nil {
		log.Fatal(err)
	}

}

func worker() error {
	for path := range jobs {
		// Carregar e salvar dados dos arquivos
		if err := service.LoadAndSaveTrades(path); err != nil {
			log.Fatal(err)
			return err
		} else {
			log.Printf("Arquivo %s processado com sucesso", path)
		}
	}
	return nil
}
