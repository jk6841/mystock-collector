package mystock

import (
	"math"
	"mystock-collector/yfin"
	"time"
)

const secondsPerMinute = 60

type Stock struct {
	Symbol string
	Name   string
}

func (stock Stock) GetData(start time.Time, end time.Time) []Candle {
	yfinResponse := yfin.Call(stock.Symbol, start, end)

	timestamp := yfinResponse.Chart.Result[0].Timestamp
	quote := yfinResponse.Chart.Result[0].Indicators.Quote[0]
	length := len(timestamp)
	candles := make([]Candle, length)
	for i := 0; i < length; i++ {
		open := rountTo2Digits(quote.Open[i])
		close := rountTo2Digits(quote.Close[i])
		high := rountTo2Digits(quote.High[i])
		low := rountTo2Digits(quote.Low[i])
		start := timestamp[i]
		candle := Candle{
			Stock: stock,
			Open:  open,
			Close: close,
			High:  high,
			Low:   low,
			Start: start,
			End:   start + secondsPerMinute,
		}
		candles[i] = candle
	}
	return candles
}

func rountTo2Digits(number float64) float64 {
	return math.Round(number*100) / 100
}
