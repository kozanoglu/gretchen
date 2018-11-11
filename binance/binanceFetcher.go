package binance

import (
	"gretchen/utils"
	"log"
	"strconv"
	"time"

	talib "github.com/markcheno/go-talib"
)

func Loop(period time.Duration, results chan<- utils.TickerList) {
	symbolsMap := GetSymbols()

	for {
		tickers := Get24HTickers()
		var tickerMap map[string]utils.Ticker
		tickerMap = make(map[string]utils.Ticker)

		for _, binanceTicker := range tickers {
			if symbolsMap[binanceTicker.Symbol].Status != "TRADING" {
				continue
			}

			klines := getCandlesForSymbol(binanceTicker.Symbol)
			rsiArray := talib.Rsi(getCloseValues(klines), 14)

			var ticker utils.Ticker
			ticker.Symbol = binanceTicker.Symbol
			ticker.Rsi = rsiArray[(len(rsiArray) - utils.Min(len(rsiArray), 7)):] // last N elements
			ticker.Price = strconv.FormatFloat(binanceTicker.LastPrice, 'f', -1, 64)
			ticker.Volume = strconv.FormatFloat(binanceTicker.Volume, 'f', -1, 64)
			ticker.QuoteVolume = strconv.FormatFloat(binanceTicker.QuoteVolume, 'f', -1, 64)
			ticker.QuoteCurrency = symbolsMap[binanceTicker.Symbol].QuoteAsset
			ticker.PriceChange1H = utils.PercentageDiff(binanceTicker.LastPrice, klines[len(klines)-2].Close)
			if len(klines) >= 5 {
				ticker.PriceChange4H = utils.PercentageDiff(binanceTicker.LastPrice, klines[len(klines)-5].Close)
			}
			if len(klines) >= 25 {
				ticker.PriceChange24H = utils.PercentageDiff(binanceTicker.LastPrice, klines[len(klines)-25].Close)
			}
			tickerMap[ticker.Symbol] = ticker
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
