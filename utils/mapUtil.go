package utils

import (
	"fmt"
	"sort"
)

// Deprecated
func SortTickerMapByRSIValues(inputMap map[string]Ticker) TickerList {
	p := make(TickerList, len(inputMap))

	i := 0
	for _, v := range inputMap {
		p[i] = v
		i++
	}

	sort.Sort(p)

	//	fmt.Printf("Post-sorted: ")
	//	for _, k := range p {
	//		fmt.Println(k.Key, ": ", k.Value)
	//	}

	return p
}

func PrintTickerList(list TickerList) {

	for _, p := range list {
		fmt.Println(p)
	}
}

type Ticker struct {
	Symbol         string
	Price          string
	Volume         string
	QuoteVolume    string
	Rsi1H          []float64
	Rsi1D          []float64
	QuoteCurrency  string
	PriceChange1H  float64
	PriceChange4H  float64
	PriceChange24H float64
}

type MarketData struct {
	Tickers TickerList
	Market  string
}

type TickerList []Ticker

func (p TickerList) Len() int      { return len(p) }
func (p TickerList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p TickerList) Less(i, j int) bool {
	return p[i].Rsi1H[len(p[i].Rsi1H)-1] < p[j].Rsi1H[len(p[j].Rsi1H)-1]
}
