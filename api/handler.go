package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/igorfarodrigues/api-cotacoes-b3/service"
	"log"
	"net/http"
	"regexp"
)

func GetTradeData(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")
	date := chi.URLParam(r, "date")

	if ticker == "" {
		log.Printf("Ticker não pode ser vazio")
		http.Error(w, "Ticker não pode ser vazio", http.StatusBadRequest)
		return
	}

	match, err := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, date)
	if err != nil {
		log.Printf("Erro ao validar data: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if !match {
		log.Printf("Formato de data inválido: %s", date)
		http.Error(w, "Formato de data inválido, use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	data, err := service.GetTradeData(ticker, date)
	if err != nil {
		log.Printf("Erro ao buscar cotações: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
