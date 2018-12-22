package kucoin

import (
	"gretchen/utils"
	"log"
	"math"
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

			if len(hourlyKlines.Close) >= 2 {
				ticker.PriceChange1H = utils.PercentageDiff(hourlyKlines.Close[len(hourlyKlines.Close)-1], hourlyKlines.Open[len(hourlyKlines.Open)-1])
			}
			if len(hourlyKlines.Close) >= 5 {
				ticker.PriceChange4H = utils.PercentageDiff(hourlyKlines.Open[len(hourlyKlines.Open)-5], hourlyKlines.Close[len(hourlyKlines.Close)-5])
			}

			ticker.PriceChange24H = math.Round(kucoinTicker.ChangeRate*100*100) / 100

			if len(hourlyKlines.Close) > 14 {
				rsiArray := talib.Rsi(hourlyKlines.Close, 14)
				ticker.Rsi1H = rsiArray[(len(rsiArray) - utils.Min(len(rsiArray), 7)):] // last N elements
			}

			if len(dailyKlines.Close) > 14 {
				rsiArray := talib.Rsi(dailyKlines.Close, 14)
				ticker.Rsi1D = rsiArray[(len(rsiArray) - utils.Min(len(rsiArray), 7)):] // last N elements
			}

			marketMap[ticker.QuoteCurrency] = append(marketMap[ticker.QuoteCurrency], ticker)
		}

		log.Println("Kucoin result are fetched")
		results <- marketMap

		time.Sleep(period * time.Second)
	}
}
