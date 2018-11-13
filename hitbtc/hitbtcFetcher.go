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
			klines := getCandlesForSymbol(hitbtcTicker.Symbol)

			if len(klines) <= 14 {
				continue
			}

			rsiArray := talib.Rsi(getCloseValues(klines), 14)

			var ticker utils.Ticker
			ticker.Symbol = hitbtcTicker.Symbol
			ticker.Rsi = rsiArray[(len(rsiArray) - utils.Min(len(rsiArray), 7)):] // last N elements
			ticker.Price = hitbtcTicker.Last
			ticker.Volume = hitbtcTicker.Volume
			ticker.QuoteVolume = hitbtcTicker.VolumeQuote
			ticker.QuoteCurrency = symbolsMap[hitbtcTicker.Symbol].QuoteCurrency

			lastPrice, err := strconv.ParseFloat(hitbtcTicker.Last, 64)
			if err != nil {
				lastPrice = 0.0
			}

			ticker.PriceChange1H = utils.PercentageDiff(lastPrice, klines[len(klines)-2].Close)
			if len(klines) >= 5 {
				ticker.PriceChange4H = utils.PercentageDiff(lastPrice, klines[len(klines)-5].Close)
			}
			if len(klines) >= 25 {
				ticker.PriceChange24H = utils.PercentageDiff(lastPrice, klines[len(klines)-25].Close)
			}

			marketMap[ticker.QuoteCurrency] = append(marketMap[ticker.QuoteCurrency], ticker)
		}

		log.Println("HitBTC result are fetched")
		results <- marketMap

		time.Sleep(period * time.Second)
	}
}

func getCandlesForSymbol(symbol string) []HitBTCCandle {
	klines := GetCandles(symbol)
	return klines
}

func getCloseValues(klines []HitBTCCandle) []float64 {
	var result = make([]float64, len(klines))
	for i, kline := range klines {
		result[i] = kline.Close
	}
	return result
}
