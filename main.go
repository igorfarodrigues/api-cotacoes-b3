package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/igorfarodrigues/api-cotacoes-b3/api"
	"github.com/igorfarodrigues/api-cotacoes-b3/configs"
	"github.com/igorfarodrigues/api-cotacoes-b3/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Log the loaded configuration
	log.Printf("Loaded configuration: %+v", configs.GetDB())

	// Carregar e salvar dados dos arquivos
	if err := service.LoadAndSaveTrades(configs.GetFolderPath()); err != nil {
		log.Fatal(err)
	}
	log.Println("Loading files - success!!!")

	r := chi.NewRouter()
	r.Get("/trades", api.GetTradeData)

	port := configs.GetServerPort()
	log.Printf("Server starting on port %s", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatal(err)
	}

}
