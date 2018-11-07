package hitbtc

import (
	"gretchen/utils"
	"log"
	"strings"
	"time"

	talib "github.com/markcheno/go-talib"
)

func Loop(period time.Duration, results chan<- utils.TickerList) {
	for {
		hitbtcTickers := Get24HTickers()
		var tickerMap map[string]utils.Ticker
		tickerMap = make(map[string]utils.Ticker)

		for _, hitbtcTicker := range hitbtcTickers {
			if strings.HasSuffix(hitbtcTicker.Symbol, "BTC") {
				klines := getCandlesForSymbol(hitbtcTicker.Symbol)
				if len(klines) > 14 {
					rsiArray := talib.Rsi(getCloseValues(klines), 14)

					var ticker utils.Ticker
					ticker.Symbol = hitbtcTicker.Symbol
					ticker.Rsi = rsiArray[(len(rsiArray) - utils.Min(len(rsiArray), 7)):] // last N elements
					ticker.Price = hitbtcTicker.Last
					ticker.Volume = hitbtcTicker.Volume

					tickerMap[ticker.Symbol] = ticker
				}
			}
		}
		log.Println("HitBTC result are fetched")
		results <- (utils.SortTickerMapByRSIValues(tickerMap))

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
