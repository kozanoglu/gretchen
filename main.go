package main

import (
	"fmt"
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
	//go binance.Loop(5)

	hitbtcResults := make(chan utils.TickerList)
	go hitbtc.Loop(60, hitbtcResults)

	go web.Start(hitbtcResults)

	for {
		time.Sleep(10 * time.Second)
	}

	fmt.Scanln()
	log.Println("Exiting the application")
}
