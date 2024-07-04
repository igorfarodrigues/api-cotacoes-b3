package main

import (
	"database/sql"
	"github.com/igorfarodrigues/api-cotacoes-b3/api"
	"github.com/igorfarodrigues/api-cotacoes-b3/repository"
	"github.com/igorfarodrigues/api-cotacoes-b3/service"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres", "user=youruser password=yourpassword dbname=yourdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	tradeRepo := repository.NewTradeRepository(db)
	tradeService := service.NewTradeService(tradeRepo)
	handler := api.NewHandler(tradeService)

	http.HandleFunc("/trades", handler.GetTradeData)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
