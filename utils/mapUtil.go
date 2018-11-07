package utils

import (
	"fmt"
	"sort"
)

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
	Symbol string
	Price  string
	Volume string
	Rsi    []float64
}

type TickerList []Ticker

func (p TickerList) Len() int           { return len(p) }
func (p TickerList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p TickerList) Less(i, j int) bool { return p[i].Rsi[len(p[i].Rsi)-1] < p[j].Rsi[len(p[j].Rsi)-1] }
