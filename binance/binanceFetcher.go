package binance

import (
	"fmt"
	"gretchen/utils"
	"strings"
	"time"

	talib "github.com/markcheno/go-talib"
)

func Loop(period time.Duration) {
	for {
		tickers := Get24HTickers()
		var rsiMap map[string]float64
		rsiMap = make(map[string]float64)

		for _, ticker := range tickers {
			if strings.HasSuffix(ticker.Symbol, "ETH") {
				klines := getCandlesForSymbol(ticker.Symbol)
				rsi := talib.Rsi(getCloseValues(klines), 14)
				rsiMap[ticker.Symbol] = rsi[len(rsi)-1]
				//fmt.Println("1H RSI for", ticker.Symbol, " is: ", rsi[len(rsi)-1])
			}
		}

		utils.PrintPairList(utils.SortMapByValues(rsiMap))
		fmt.Println("*********")

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
