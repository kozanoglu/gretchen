package main

import (
	"fmt"
	"gretchen/hitbtc"
	"gretchen/utils"
	"gretchen/web"
)

const SleepInterval = 5 * 1000000000

func main() {

	fmt.Println("Starting the application...")

	fmt.Println("Main loop thread started...")
	//go idex.Loop(5)
	//go binance.Loop(5)

	hitbtcResults := make(chan utils.PairList)
	go hitbtc.Loop(10, hitbtcResults)

	go web.Start(hitbtcResults)

	/*
		res := <-hitbtcResults
		utils.PrintPairList(res)*/

	fmt.Scanln()
	fmt.Println("Exit")
}

/*
func dispatchMessages() {
	for {
        msg := <- ch
        for _, worker := workers {
            worker.source <- msg
        }
    }
}*/
