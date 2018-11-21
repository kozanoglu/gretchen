package hitbtc

import (
	"gretchen/utils"
	"log"
	"strconv"
	"time"

	talib "github.com/markcheno/go-talib"
)

func Loop(period time.Duration, results chan<- map[string][]utils.Ticker) {
	symbolsMap := GetSymbols()

	for {
		hitbtcTickers := Get24HTickers()
		marketMap := map[string][]utils.Ticker{}

		for _, hitbtcTicker := range hitbtcTickers {
			hourlyKlines := GetHourlyCandles(hitbtcTicker.Symbol)
			dailyKlines := GetDailyCandles(hitbtcTicker.Symbol)

			var ticker utils.Ticker
			ticker.Symbol = hitbtcTicker.Symbol
			ticker.Price = hitbtcTicker.Last
			ticker.Volume = hitbtcTicker.Volume
			ticker.QuoteVolume = hitbtcTicker.VolumeQuote
			ticker.QuoteCurrency = symbolsMap[hitbtcTicker.Symbol].QuoteCurrency

			lastPrice, err := strconv.ParseFloat(hitbtcTicker.Last, 64)
			if err != nil {
				lastPrice = 0.0
			}

			ticker.PriceChange1H = utils.PercentageDiff(lastPrice, hourlyKlines[len(hourlyKlines)-2].Close)
			if len(hourlyKlines) >= 5 {
				ticker.PriceChange4H = utils.PercentageDiff(lastPrice, hourlyKlines[len(hourlyKlines)-5].Close)
			}
			if len(hourlyKlines) >= 25 {
				ticker.PriceChange24H = utils.PercentageDiff(lastPrice, hourlyKlines[len(hourlyKlines)-25].Close)
			}

			if len(hourlyKlines) > 14 {
				rsiArray := talib.Rsi(getCloseValues(hourlyKlines), 14)
				ticker.Rsi = rsiArray[(len(rsiArray) - utils.Min(len(rsiArray), 7)):] // last N elements
			}

			if len(dailyKlines) > 14 {
				rsiArray := talib.Rsi(getCloseValues(dailyKlines), 14)
				ticker.Rsi1D = rsiArray[(len(rsiArray) - utils.Min(len(rsiArray), 7)):] // last N elements
			}

			marketMap[ticker.QuoteCurrency] = append(marketMap[ticker.QuoteCurrency], ticker)
		}

		log.Println("HitBTC result are fetched")
		results <- marketMap

		time.Sleep(period * time.Second)
	}
}

func getCloseValues(klines []HitBTCCandle) []float64 {
	var result = make([]float64, len(klines))
	for i, kline := range klines {
		result[i] = kline.Close
	}
	return result
}
