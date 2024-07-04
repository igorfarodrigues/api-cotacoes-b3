package api

import (
    "github.com/igorfarodrigues/api-cotacoes-b3/service"
    "encoding/json"
    "net/http"
)

type Handler struct {
    TradeService *service.TradeService
}

func NewHandler(service *service.TradeService) *Handler {
    return &Handler{TradeService: service}
}

func (handler *Handler) GetTradeData(w http.ResponseWriter, r *http.Request) {
    ticker := r.URL.Query().Get("ticker")
    date := r.URL.Query().Get("date")

    if ticker == "" {
        http.Error(w, "Ticker is required", http.StatusBadRequest)
        return
    }

    data, err := handler.TradeService.GetTradeData(ticker, date)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(data)
}
