package service

import (
    "github.com/igorfarodrigues/api-cotacoes-b3/models"
    "github.com/igorfarodrigues/api-cotacoes-b3/repository"
)

type TradeService struct {
    Repo *repository.TradeRepository
}

func NewTradeService(repo *repository.TradeRepository) *TradeService {
    return &TradeService{Repo: repo}
}

func (service *TradeService) SaveTrade(trade *models.Trade) error {
    return service.Repo.SaveTrade(trade)
}

func (service *TradeService) GetTradeData(ticker string, date string) (map[string]interface{}, error) {
    maxRangeValue, err := service.Repo.GetMaxRangeValue(ticker, date)
    if err != nil {
        return nil, err
    }

    maxDailyVolume, err := service.Repo.GetMaxDailyVolume(ticker, date)
    if err != nil {
        return nil, err
    }

    result := map[string]interface{}{
        "ticker":           ticker,
        "max_range_value":  maxRangeValue,
        "max_daily_volume": maxDailyVolume,
    }
    return result, nil
}
