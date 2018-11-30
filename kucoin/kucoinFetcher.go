package kucoin

import (
	"gretchen/utils"
	"log"
	"strconv"
	"time"

	talib "github.com/markcheno/go-talib"
)

func Loop(period time.Duration, results chan<- map[string][]utils.Ticker) {
	/*	markets := GetMarkets()
		fmt.Println(markets)*/

	for {
		kucoinTickers := Get24HTickers()
		marketMap := map[string][]utils.Ticker{}

		for _, kucoinTicker := range kucoinTickers {
			hourlyKlines := GetHourlyCandles(kucoinTicker.Symbol)
			dailyKlines := GetDailyCandles(kucoinTicker.Symbol)

			var ticker utils.Ticker
			ticker.Symbol = kucoinTicker.Symbol
			ticker.Price = strconv.FormatFloat(kucoinTicker.LastDealPrice, 'f', -1, 64)
			ticker.Volume = strconv.FormatFloat(kucoinTicker.Vol, 'f', -1, 64)
			ticker.QuoteVolume = strconv.FormatFloat(kucoinTicker.VolValue, 'f', -1, 64)
			ticker.QuoteCurrency = kucoinTicker.CoinTypePair

			if len(hourlyKlines) >= 2 {
				ticker.PriceChange1H = utils.PercentageDiff(kucoinTicker.LastDealPrice, hourlyKlines[len(hourlyKlines)-2].Close)
			}
			if len(hourlyKlines) >= 5 {
				ticker.PriceChange4H = utils.PercentageDiff(kucoinTicker.LastDealPrice, hourlyKlines[len(hourlyKlines)-5].Close)
			}
			if len(hourlyKlines) >= 25 {
				ticker.PriceChange24H = kucoinTicker.ChangeRate * 100 // todo fix decimals
			}

			if len(hourlyKlines) > 14 {
				rsiArray := talib.Rsi(getCloseValues(hourlyKlines), 14)
				ticker.Rsi1H = rsiArray[(len(rsiArray) - utils.Min(len(rsiArray), 7)):] // last N elements
			}

			if len(dailyKlines) > 14 {
				rsiArray := talib.Rsi(getCloseValues(dailyKlines), 14)
				ticker.Rsi1D = rsiArray[(len(rsiArray) - utils.Min(len(rsiArray), 7)):] // last N elements
			}

			marketMap[ticker.QuoteCurrency] = append(marketMap[ticker.QuoteCurrency], ticker)
		}

		log.Println("Kucoin result are fetched")
		results <- marketMap

		time.Sleep(period * time.Second)
	}
}

func getCloseValues(klines []*KucoinCandle) []float64 {
	var result = make([]float64, len(klines))
	for i, kline := range klines {
		result[i] = kline.Close
	}
	return result
}
