package utils

import (
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	//	fmt.Println("response Status:", resp.Status)
	//	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
