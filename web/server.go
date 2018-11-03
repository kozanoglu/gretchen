package web

import (
	"gretchen/utils"
	"html/template"
	"log"
	"net/http"
	"os"
)

var hitbtcPairs utils.TickerList
var pageTemplate = template.Must(template.ParseFiles("templates/index.html"))

var handler = func(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		PageTitle: "HitBTC Results",
		Tickers:   hitbtcPairs,
	}
	pageTemplate.Execute(w, data)
}

func Start(hitbtcChannel chan utils.TickerList) {

	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT not set, using default 8000")
		port = "8000"
	}

	http.HandleFunc("/", handler)

	go http.ListenAndServe(":"+port, nil)
	log.Println("Started the web server on port 8000")

	for {
		hitbtcPairs = <-hitbtcChannel
	}
}

type PageData struct {
	PageTitle string
	Tickers   utils.TickerList
}
