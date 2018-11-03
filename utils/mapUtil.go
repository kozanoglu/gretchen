package utils

import (
	"fmt"
	"sort"
)

func SortMapByValues(inputMap map[string]float64) PairList {
	p := make(PairList, len(inputMap))

	i := 0
	for k, v := range inputMap {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	//	fmt.Printf("Post-sorted: ")
	//	for _, k := range p {
	//		fmt.Println(k.Key, ": ", k.Value)
	//	}

	return p
}

func PrintPairList(list PairList) {

	for _, p := range list {
		fmt.Println(p.Key, ": ", p.Value)
	}
}

type Pair struct {
	Key   string
	Value float64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
