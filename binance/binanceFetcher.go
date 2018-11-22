package binance

import (
	"gretchen/utils"
	"log"
	"strconv"
	"strings"
	"time"

	talib "github.com/markcheno/go-talib"
)

func Loop(period time.Duration, results chan<- map[string][]utils.Ticker) {
	symbolsMap := GetSymbols()

	for {
		tickers := Get24HTickers()
		marketMap := map[string][]utils.Ticker{}

		for _, binanceTicker := range tickers {
			if symbolsMap[binanceTicker.Symbol].Status != "TRADING" {
				continue
			}

			hourlyKlines := GetHourlyCandles(binanceTicker.Symbol)
			dailyKlines := GetDailyCandles(binanceTicker.Symbol)

			quoteAsset := symbolsMap[binanceTicker.Symbol].QuoteAsset

			var ticker utils.Ticker
			ticker.Symbol = strings.Replace(binanceTicker.Symbol, quoteAsset, "_"+quoteAsset, 1)
			ticker.Price = strconv.FormatFloat(binanceTicker.LastPrice, 'f', -1, 64)
			ticker.Volume = strconv.FormatFloat(binanceTicker.Volume, 'f', -1, 64)
			ticker.QuoteVolume = strconv.FormatFloat(binanceTicker.QuoteVolume, 'f', -1, 64)
			ticker.QuoteCurrency = quoteAsset
			ticker.PriceChange1H = utils.PercentageDiff(binanceTicker.LastPrice, hourlyKlines[len(hourlyKlines)-2].Close)
			if len(hourlyKlines) >= 5 {
				ticker.PriceChange4H = utils.PercentageDiff(binanceTicker.LastPrice, hourlyKlines[len(hourlyKlines)-5].Close)
			}
			if len(hourlyKlines) >= 25 {
				ticker.PriceChange24H = utils.PercentageDiff(binanceTicker.LastPrice, hourlyKlines[len(hourlyKlines)-25].Close)
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

		log.Println("Binance result are fetched")
		results <- marketMap

		time.Sleep(period * time.Second)
	}
}

func getCloseValues(klines []*BinanceCandle) []float64 {
	var result = make([]float64, len(klines))
	for i, kline := range klines {
		result[i] = kline.Close
	}
	return result
}
