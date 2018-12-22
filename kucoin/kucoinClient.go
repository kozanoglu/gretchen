package kucoin

import (
	"encoding/json"
	"gretchen/utils"
	"log"
	"strconv"
	"time"

	"github.com/golang/glog"
)

const KucoinEndpoint = "https://api.kucoin.com"
const MarketsAPI = "/v1/open/markets"
const TickersAPI = "/v1/market/open/symbols"
const KlinesAPI = "/v1/open/kline"
const KlinesAPI2 = "/v1/open/chart/history"

func GetMarkets() KucoinMarketResult {
	url := KucoinEndpoint + MarketsAPI
	body, err := utils.Get(url)
	if err != nil {
		glog.Error(err)
	}
	return parseMarketInfo(body)
}

func Get24HTickers() map[string]KucoinTicker {
	url := KucoinEndpoint + TickersAPI
	body, err := utils.Get(url)
	if err != nil {
		glog.Error(err)
		return nil
	}
	return parseSymbolInfo(body)
}

func GetHourlyCandles(symbol string) *KucoinKlineResult {
	//?symbol=KCS-BTC&type=1HOUR&limit=10&from=1540145736&to=1542145736
	to := time.Now().Unix()
	toMs := strconv.FormatInt(to, 10)
	from := to - 340*3600
	fromMs := strconv.FormatInt(from, 10)
	url := KucoinEndpoint + KlinesAPI2 + "?symbol=" + symbol + "&resolution=60&from=" + fromMs + "&to=" + toMs

	body, err := utils.Get(url)
	if err != nil {
		glog.Error(err)
		return nil
	}
	return parseKlinesInfo(body)
}

func GetDailyCandles(symbol string) *KucoinKlineResult {
	//?symbol=KCS-BTC&type=1DAY&limit=10&from=1540145736&to=1542145736
	to := time.Now().Unix()
	toMs := strconv.FormatInt(to, 10)
	from := to - 340*3600*24
	fromMs := strconv.FormatInt(from, 10)
	url := KucoinEndpoint + KlinesAPI2 + "?symbol=" + symbol + "&resolution=D&from=" + fromMs + "&to=" + toMs
	body, err := utils.Get(url)
	if err != nil {
		glog.Error(err)
		return nil
	}
	return parseKlinesInfo(body)
}

func parseMarketInfo(input []byte) KucoinMarketResult {
	var result KucoinMarketResult
	err := json.Unmarshal(input, &result)

	if err != nil {
		glog.Error(err)
	}

	return result
}

func parseSymbolInfo(input []byte) map[string]KucoinTicker {
	var result KucoinSymbolResult
	err := json.Unmarshal(input, &result)

	if err != nil {
		panic(err)
	}

	resultAsMap := make(map[string]KucoinTicker)
	for i := 0; i < len(result.Data); i++ {
		resultAsMap[result.Data[i].Symbol] = result.Data[i]
	}

	return resultAsMap
}

func parseKlinesInfo(input []byte) *KucoinKlineResult {
	var parsedBody KucoinKlineResult
	err := json.Unmarshal(input, &parsedBody)

	if err != nil {
		log.Panic("Failed to parse result: ", string(input), err)
	}

	return &parsedBody
}

type KucoinMarketResult struct {
	Success   bool     `json:"success"`
	Code      string   `json:"code"`
	Msg       string   `json:"msg"`
	Timestamp int64    `json:"timestamp"`
	Data      []string `json:"data"`
}

type KucoinSymbolResult struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
	Data      []KucoinTicker
}

type KucoinTicker struct {
	CoinType      string  `json:"coinType"`
	Trading       bool    `json:"trading"`
	Symbol        string  `json:"symbol"`
	LastDealPrice float64 `json:"lastDealPrice"`
	Buy           float64 `json:"buy"`
	Sell          float64 `json:"sell"`
	Change        float64 `json:"change"`
	CoinTypePair  string  `json:"coinTypePair"`
	Sort          int     `json:"sort"`
	FeeRate       float64 `json:"feeRate"`
	VolValue      float64 `json:"volValue"`
	High          float64 `json:"high"`
	Datetime      int64   `json:"datetime"`
	Vol           float64 `json:"vol"`
	Low           float64 `json:"low"`
	ChangeRate    float64 `json:"changeRate"`
}

type KucoinKlineResult struct {
	Success   bool      `json:"success"`
	Code      string    `json:"code"`
	Status    string    `json:"s"`
	Close     []float64 `json:"c"`
	Timestamp []int     `json:"t"`
	Volume    []float64 `json:"v"`
	High      []float64 `json:"h"`
	Low       []float64 `json:"l"`
	Open      []float64 `json:"o"`
}
