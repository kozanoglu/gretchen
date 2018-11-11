package binance

import (
	"encoding/json"
	"gretchen/utils"
)

const BinanceEndpoint = "https://api.binance.com"
const ExchangeInfoAPI = "/api/v1/exchangeInfo"
const DailyTickerAPI = "/api/v1/ticker/24hr"
const KlinesAPI = "/api/v1/klines"

func GetSymbols() map[string]BinanceSymbol {
	exchangeInfo := GetExchangeInfo()
	return parseSymbolInfo(exchangeInfo.Symbols)
}

func GetExchangeInfo() ExchangeInfo {
	url := BinanceEndpoint + ExchangeInfoAPI
	body := utils.Get(url)
	return parseExchangeInfo(body)
}

func Get24HTickers() []BinanceTicker {
	url := BinanceEndpoint + DailyTickerAPI
	body := utils.Get(url)
	return parseTickerInfo(body)
}

func GetCandles(symbol string) []*BinanceCandle {
	url := BinanceEndpoint + KlinesAPI + "?symbol=" + symbol + "&interval=1h&limit=336"
	body := utils.Get(url)

	//fmt.Println(string(body))
	return parseKlinesInfo(body)
}

func parseExchangeInfo(input []byte) ExchangeInfo {
	var result ExchangeInfo
	err := json.Unmarshal(input, &result)

	if err != nil {
		panic(err)
	}
	return result
}

func parseSymbolInfo(symbolArray []BinanceSymbol) map[string]BinanceSymbol {
	resultAsMap := make(map[string]BinanceSymbol)
	for i := 0; i < len(symbolArray); i++ {
		resultAsMap[symbolArray[i].Symbol] = symbolArray[i]
	}
	return resultAsMap
}

func parseTickerInfo(input []byte) []BinanceTicker {
	var result []BinanceTicker
	err := json.Unmarshal(input, &result)

	if err != nil {
		panic(err)
	}
	return result
}

func parseKlinesInfo(input []byte) []*BinanceCandle {
	var parsedBody [][]interface{}
	err := json.Unmarshal(input, &parsedBody)

	if err != nil {
		panic(err)
	}

	var result = make([]*BinanceCandle, len(parsedBody))
	for i, klineArray := range parsedBody {
		result[i] = NewCandle(klineArray)
	}
	return result
}

type BinanceSymbol struct {
	Symbol             string   `json:"symbol"`
	Status             string   `json:"status"`
	BaseAsset          string   `json:"baseAsset"`
	BaseAssetPrecision int      `json:"baseAssetPrecision"`
	QuoteAsset         string   `json:"quoteAsset"`
	QuotePrecision     int      `json:"quotePrecision"`
	OrderTypes         []string `json:"orderTypes"`
	IcebergAllowed     bool     `json:"icebergAllowed"`
	Filters            []struct {
		FilterType  string `json:"filterType"`
		MinPrice    string `json:"minPrice,omitempty"`
		MaxPrice    string `json:"maxPrice,omitempty"`
		TickSize    string `json:"tickSize,omitempty"`
		MinQty      string `json:"minQty,omitempty"`
		MaxQty      string `json:"maxQty,omitempty"`
		StepSize    string `json:"stepSize,omitempty"`
		MinNotional string `json:"minNotional,omitempty"`
	} `json:"filters"`
}

type ExchangeInfo struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"serverTime"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
	ExchangeFilters []interface{}   `json:"exchangeFilters"`
	Symbols         []BinanceSymbol `json:"symbols"`
}

type BinanceTicker struct {
	Symbol             string  `json:"symbol"`
	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`
	PrevClosePrice     float64 `json:"prevClosePrice,string"`
	LastPrice          float64 `json:"lastPrice,string"`
	LastQty            float64 `json:"lastQty,string"`
	BidPrice           float64 `json:"bidPrice,string"`
	AskPrice           float64 `json:"askPrice,string"`
	OpenPrice          float64 `json:"openPrice,string"`
	HighPrice          float64 `json:"highPrice,string"`
	LowPrice           float64 `json:"lowPrice,string"`
	Volume             float64 `json:"volume,string"`
	QuoteVolume        float64 `json:"quoteVolume,string"`
	OpenTime           int64   `json:"openTime"`
	CloseTime          int64   `json:"closeTime"`
	FirstID            int     `json:"firstId"`
	LastID             int     `json:"lastId"`
	Count              int     `json:"count"`
}

type BinanceCandle struct {
	OpenTime                 int64
	Open                     float64
	High                     float64
	Low                      float64
	Close                    float64
	Volume                   float64
	CloseTime                int64
	QuoteVolume              float64
	NumberOfTrades           int32
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
	Ignore                   float64
}

func NewCandle(input []interface{}) *BinanceCandle {
	k := new(BinanceCandle)
	k.OpenTime = int64(input[0].(float64))
	k.Open, _ = utils.GetFloat(input[1])
	k.High, _ = utils.GetFloat(input[2])
	k.Low, _ = utils.GetFloat(input[3])
	k.Close, _ = utils.GetFloat(input[4])
	k.Volume, _ = utils.GetFloat(input[5])
	k.CloseTime = int64(input[6].(float64))
	k.QuoteVolume, _ = utils.GetFloat(input[7])
	k.NumberOfTrades = int32(input[8].(float64))
	k.TakerBuyBaseAssetVolume, _ = utils.GetFloat(input[9])
	k.TakerBuyQuoteAssetVolume, _ = utils.GetFloat(input[10])
	k.Ignore, _ = utils.GetFloat(input[11])
	return k
}
