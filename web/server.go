package web

import (
	"fmt"
	"gretchen/utils"
	"html/template"
	"net/http"
)

var hitbtcPairs utils.PairList
var pageTemplate = template.Must(template.ParseFiles("gretchen/web/main.html"))

var handler = func(w http.ResponseWriter, r *http.Request) {
	data := TodoPageData{
		PageTitle: "HitBTC Results",
		Pairs:     hitbtcPairs,
	}
	pageTemplate.Execute(w, data)
}

func Start(hitbtcChannel chan utils.PairList) {

	http.HandleFunc("/", handler)

	go http.ListenAndServe(":80", nil)
	fmt.Println("Started the web server on port 80")

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
