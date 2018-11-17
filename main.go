package main

import (
	"fmt"
	"gretchen/binance"
	"gretchen/hitbtc"
	"gretchen/kucoin"
	"gretchen/utils"
	"gretchen/web"
	"log"
	"time"
)

const channelTimeout int = 5

func main() {

	log.Println("Starting the application...")

	log.Println("Main loop thread started...")
	//go idex.Loop(5)

	binanceResults := make(chan map[string][]utils.Ticker, channelTimeout)
	hitbtcResults := make(chan map[string][]utils.Ticker, channelTimeout)
	kucoinResults := make(chan map[string][]utils.Ticker, channelTimeout)

	go binance.Loop(60, binanceResults)
	go hitbtc.Loop(60, hitbtcResults)
	go kucoin.Loop(60, kucoinResults)
	go web.Start(binanceResults, hitbtcResults, kucoinResults)

	for {
		time.Sleep(10 * time.Second)
	}

	fmt.Scanln()
	log.Println("Exiting the application")
}
