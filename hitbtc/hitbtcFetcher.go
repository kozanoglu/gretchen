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
		tickers := Get24HTickers()
		var rsiMap map[string]utils.Ticker
		rsiMap = make(map[string]utils.Ticker)

		for _, ticker := range tickers {
			if strings.HasSuffix(ticker.Symbol, "BTC") {
				klines := getCandlesForSymbol(ticker.Symbol)
				if len(klines) > 14 {
					rsi := talib.Rsi(getCloseValues(klines), 14)

					var unifiedTicker utils.Ticker
					unifiedTicker.Symbol = ticker.Symbol
					unifiedTicker.Rsi = rsi[len(rsi)-1]
					unifiedTicker.Price = ticker.Last
					unifiedTicker.Volume = ticker.Volume

					rsiMap[ticker.Symbol] = unifiedTicker
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
