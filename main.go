package main

import (
	"fmt"
	"gretchen/binance"
	"gretchen/hitbtc"
	"gretchen/utils"
	"gretchen/web"
	"log"
	"time"
)

func main() {

	log.Println("Starting the application...")

	log.Println("Main loop thread started...")
	//go idex.Loop(5)

	binanceResults := make(chan utils.TickerList)
	hitbtcResults := make(chan utils.TickerList)

	go binance.Loop(60, binanceResults)
	go hitbtc.Loop(60, hitbtcResults)
	go web.Start(binanceResults, hitbtcResults)

	for {
		time.Sleep(10 * time.Second)
	}

	fmt.Scanln()
	log.Println("Exiting the application")
}
