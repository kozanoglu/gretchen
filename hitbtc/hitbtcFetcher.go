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
		var rsiMap map[string]utils.Ticker
		rsiMap = make(map[string]utils.Ticker)

		for _, hitbtcTicker := range hitbtcTickers {
			if strings.HasSuffix(hitbtcTicker.Symbol, "BTC") {
				klines := getCandlesForSymbol(hitbtcTicker.Symbol)
				if len(klines) > 14 {
					rsi := talib.Rsi(getCloseValues(klines), 14)

					var ticker utils.Ticker
					ticker.Symbol = hitbtcTicker.Symbol
					ticker.Rsi = rsi[len(rsi)-1]
					ticker.Price = hitbtcTicker.Last
					ticker.Volume = hitbtcTicker.Volume

					rsiMap[ticker.Symbol] = ticker
				}
				//fmt.Println("1H RSI for", ticker.Symbol, " is: ", rsi[len(rsi)-1])
			}
		}
		log.Println("HitBTC result are fetched")

		//	utils.PrintPairList(utils.SortMapByValues(rsiMap))
		results <- (utils.SortMapByValues(rsiMap))

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
