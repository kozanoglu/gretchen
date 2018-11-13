package web

import (
	"fmt"
	"gretchen/utils"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var binancePairs map[string][]utils.Ticker
var hitbtcPairs map[string][]utils.Ticker

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
	"htmlSafe": func(html string) template.HTML {
		return template.HTML(html)
	},
}

//var pageTemplate = template.Must(template.New("main").Funcs(funcMap).ParseGlob("static/*.html"))

func Start(binanceChannel chan map[string][]utils.Ticker, hitbtcChannel chan map[string][]utils.Ticker) {

	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT not set, using default 8000")
		port = "8000"
	}

	router := gin.Default()
	router.SetFuncMap(funcMap)
	router.Static("/static", "./static")
	router.LoadHTMLGlob("static/*.html")
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/binance")
	}).GET("/binance", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"data": binancePairs,
		})
	}).GET("/hitbtc", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"data": hitbtcPairs,
		})
	})

	go router.Run(":8000")
	log.Println("Started the web server on port ", port)

	for {
		hitbtcPairs = <-hitbtcChannel
		binancePairs = <-binanceChannel
	}
}
