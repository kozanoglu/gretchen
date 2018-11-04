package binance

import (
	"gretchen/utils"
	"log"
	"strconv"
	"strings"
	"time"

	talib "github.com/markcheno/go-talib"
)

func Loop(period time.Duration, results chan<- utils.TickerList) {
	for {
		tickers := Get24HTickers()
		var tickerMap map[string]utils.Ticker
		tickerMap = make(map[string]utils.Ticker)

		for _, binanceTicker := range tickers {
			if strings.HasSuffix(binanceTicker.Symbol, "BTC") {
				klines := getCandlesForSymbol(binanceTicker.Symbol)
				rsi := talib.Rsi(getCloseValues(klines), 14)

				var ticker utils.Ticker
				ticker.Symbol = binanceTicker.Symbol
				ticker.Rsi = rsi[len(rsi)-1]
				ticker.Price = strconv.FormatFloat(binanceTicker.LastPrice, 'f', -1, 64)
				ticker.Volume = strconv.FormatFloat(binanceTicker.Volume, 'f', -1, 64)

				tickerMap[ticker.Symbol] = ticker
			}
		}

		log.Println("Binance result are fetched")
		results <- (utils.SortTickerMapByRSIValues(tickerMap))

		time.Sleep(period * time.Second)
	}
}

func getCandlesForSymbol(symbol string) []*BinanceCandle {
	klines := GetCandles(symbol)
	return klines
}

func getCloseValues(klines []*BinanceCandle) []float64 {
	var result = make([]float64, len(klines))
	for i, kline := range klines {
		result[i] = kline.Close
	}
	return result
}
