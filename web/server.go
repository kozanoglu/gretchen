package web

import (
	"gretchen/utils"
	"html/template"
	"log"
	"net/http"
	"os"
)

var hitbtcPairs utils.PairList
var pageTemplate = template.Must(template.ParseFiles("templates/index.html"))

var handler = func(w http.ResponseWriter, r *http.Request) {
	data := TodoPageData{
		PageTitle: "HitBTC Results",
		Pairs:     hitbtcPairs,
	}
	pageTemplate.Execute(w, data)
}

func Start(hitbtcChannel chan utils.PairList) {

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

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Pairs     utils.PairList
}
