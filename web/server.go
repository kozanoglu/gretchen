package web

import (
	"fmt"
	"gretchen/utils"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var binancePairs utils.TickerList
var hitbtcPairs utils.TickerList

var funcMap = template.FuncMap{
	"FormatFloat": func(f float64) string { return fmt.Sprintf("%.2f", f) },
	"LastElem":    func(f []float64) string { return fmt.Sprintf("%.2f", f[len(f)-1]) },
	"ToJsArrayFunction": func(f []float64) template.JS {
		arr := "["
		for _, v := range f {
			arr += fmt.Sprintf("%.2f", v) + ","
		}
		arr = strings.TrimSuffix(arr, ",") + "]"
		result := "showChart(" + arr + ")"
		return template.JS(result)
	},
}

var handler = func(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		HitBtcTickers:  hitbtcPairs,
		BinanceTickers: binancePairs,
	}
	pageTemplate.ExecuteTemplate(w, "index.html", data)
}

var pageTemplate = template.Must(template.New("main").Funcs(funcMap).ParseGlob("static/*.html"))

func Start(binanceChannel chan utils.TickerList, hitbtcChannel chan utils.TickerList) {

	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT not set, using default 8000")
		port = "8000"
	}

	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	go http.ListenAndServe(":"+port, nil)
	log.Println("Started the web server on port ", port)

	for {
		hitbtcPairs = <-hitbtcChannel
		binancePairs = <-binanceChannel
	}
}

type PageData struct {
	HitBtcTickers  utils.TickerList
	BinanceTickers utils.TickerList
}
