package yfin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// https://query2.finance.yahoo.com/v8/finance/chart/AAPL?period1=1683747000&period2=1683747300

const BaseUri = "https://query2.finance.yahoo.com/v8/finance/chart/%s?period1=%d&period2=%d"

func Call(ticker string, start time.Time, end time.Time) YahooFinanceResponse {
	client := &http.Client{}

	url := fmt.Sprintf(BaseUri, ticker, start.Unix(), end.Unix())

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Read the response body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var yahooFinanceResponse YahooFinanceResponse
	err = json.Unmarshal(body, &yahooFinanceResponse)
	if err != nil {
		panic(err)
	}

	return yahooFinanceResponse
}
