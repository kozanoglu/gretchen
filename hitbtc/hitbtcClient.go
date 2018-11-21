package hitbtc

import (
	"encoding/json"
	"gretchen/utils"
	"time"

	"github.com/golang/glog"
)

const HitBTCEndpoint = "https://api.hitbtc.com"
const SymbolsAPI = "/api/2/public/symbol"
const DailyTickerAPI = "/api/2/public/ticker"
const KlinesAPI = "/api/2/public/candles/"

func GetSymbols() map[string]HitBTCSymbol {
	url := HitBTCEndpoint + SymbolsAPI
	body, err := utils.Get(url)
	if err != nil {
		glog.Error(err)
		return nil
	}
	return parseSymbolInfo(body)
}

func Get24HTickers() []HitBTCTicker {
	url := HitBTCEndpoint + DailyTickerAPI
	body, err := utils.Get(url)
	if err != nil {
		glog.Error(err)
		return nil
	}

	return parseTickerInfo(body)
}

func GetHourlyCandles(symbol string) []HitBTCCandle {
	url := HitBTCEndpoint + KlinesAPI + symbol + "?period=H1&limit=336"
	body, err := utils.Get(url)
	if err != nil {
		glog.Error(err)
		return nil
	}

	//fmt.Println(string(body))
	return parseKlinesInfo(body)
}

func GetDailyCandles(symbol string) []HitBTCCandle {
	url := HitBTCEndpoint + KlinesAPI + symbol + "?period=D1&limit=336"
	body, err := utils.Get(url)
	if err != nil {
		glog.Error(err)
		return nil
	}

	//fmt.Println(string(body))
	return parseKlinesInfo(body)
}

func parseSymbolInfo(input []byte) map[string]HitBTCSymbol {
	var result []HitBTCSymbol
	err := json.Unmarshal(input, &result)

	if err != nil {
		glog.Error(err)
		return nil
	}

	resultAsMap := make(map[string]HitBTCSymbol)
	for i := 0; i < len(result); i++ {
		resultAsMap[result[i].ID] = result[i]
	}

	return resultAsMap
}

func parseTickerInfo(input []byte) []HitBTCTicker {
	var result []HitBTCTicker
	err := json.Unmarshal(input, &result)

	if err != nil {
		glog.Error(err)
	}
	return result
}

func parseKlinesInfo(input []byte) []HitBTCCandle {
	var result []HitBTCCandle

	err := json.Unmarshal(input, &result)

	if err != nil {
		glog.Error(err)
		return nil
	}

	return result
}

type HitBTCSymbol struct {
	ID                   string `json:"id"`
	BaseCurrency         string `json:"baseCurrency"`
	QuoteCurrency        string `json:"quoteCurrency"`
	QuantityIncrement    string `json:"quantityIncrement"`
	TickSize             string `json:"tickSize"`
	TakeLiquidityRate    string `json:"takeLiquidityRate"`
	ProvideLiquidityRate string `json:"provideLiquidityRate"`
	FeeCurrency          string `json:"feeCurrency"`
}

type HitBTCTicker struct {
	Ask         string    `json:"ask"`
	Bid         string    `json:"bid"`
	Last        string    `json:"last"`
	Open        string    `json:"open"`
	Low         string    `json:"low"`
	High        string    `json:"high"`
	Volume      string    `json:"volume"`
	VolumeQuote string    `json:"volumeQuote"`
	Timestamp   time.Time `json:"timestamp"`
	Symbol      string    `json:"symbol"`
}

type HitBTCCandle struct {
	Timestamp   time.Time `json:"timestamp"`
	Open        float64   `json:"open,string"`
	Close       float64   `json:"close,string"`
	Min         float64   `json:"min,string"`
	Max         float64   `json:"max,string"`
	Volume      float64   `json:"volume,string"`
	VolumeQuote float64   `json:"volumeQuote,string"`
}
