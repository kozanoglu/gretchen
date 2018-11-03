package idex

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func Loop(period time.Duration) {
	for {
		result := GetEtherMarkets()
		//ticker := result["ETH_CHX"]
		//fmt.Println("ETH_CHX: " + ticker.String())

		for k, v := range result {
			baseVolume, err := strconv.ParseFloat(v.BaseVolume, 64)
			if err != nil {
				panic(err)
			}

			percentChange, err := strconv.ParseFloat(v.PercentChange, 64)
			if err != nil {
				panic(err)
			}

			if baseVolume > 5.0 && math.Abs(percentChange) > 15 {
				fmt.Println(k, ": ", v)
			}
		}

		time.Sleep(period * time.Second)
	}
}
