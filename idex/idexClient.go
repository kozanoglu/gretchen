package idex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const IdexEndpoint = "https://api.idex.market/"

type EthMarket struct {
	Last          string `json:"last"`
	High          string `json:"high"`
	Low           string `json:"low"`
	LowestAsk     string `json:"lowestAsk"`
	HighestBid    string `json:"highestBid"`
	PercentChange string `json:"percentChange"`
	BaseVolume    string `json:"baseVolume"`
	QuoteVolume   string `json:"quoteVolume"`
}

func GetEtherMarkets() map[string]*EthMarket {
	returnTickerUrl := IdexEndpoint + "returnTicker"
	jsonStr := []byte(`{}`)
	req, err := http.NewRequest("POST", returnTickerUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//	fmt.Println("response Status:", resp.Status)
	//	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	//resultStr := string(body)
	//fmt.Println("response Body:", resultStr)
	return parse(body)
}

func parse(input []byte) map[string]*EthMarket {

	var objmap map[string]*EthMarket
	err := json.Unmarshal(input, &objmap)

	if err != nil {
		panic(err)
	}

	//ticker := objmap["ETH_CHX"]
	//fmt.Println(ticker.BaseVolume + "  " + ticker.Last)
	return objmap
}

func (m EthMarket) String() string {
	message := "[last: %s, high: %s, low: %s, lowestAsk: %s, highestBid: %s, percentChange: %s, baseVol: %s,quoteVol: %s]"
	return fmt.Sprintf(message, m.Last, m.High, m.Low, m.LowestAsk, m.HighestBid, m.PercentChange, m.BaseVolume, m.QuoteVolume)
}
