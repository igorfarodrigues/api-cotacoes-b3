package service

import (
	"github.com/igorfarodrigues/api-cotacoes-b3/repository"
	"github.com/igorfarodrigues/api-cotacoes-b3/utils"
	"os"
	"path/filepath"
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

func LoadAndSaveTrades(directory string) error {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			trades, err := utils.ReadTradesFromFile(path)
			if err != nil {
				return err
			}
			for _, trade := range trades {
				if _, err := repository.SaveTrade(trade); err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}
