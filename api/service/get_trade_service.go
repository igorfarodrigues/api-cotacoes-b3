package service

import (
	"github.com/igorfarodrigues/api-cotacoes-b3/api/repository"
)

func GetTradeData(ticker string, date string) (map[string]interface{}, error) {
	maxRangeValue, err := repository.GetMaxRangeValue(ticker, date)
	if err != nil {
		return nil, err
	}

	maxDailyVolume, err := repository.GetMaxDailyVolume(ticker, date)
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
