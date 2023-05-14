package yfin

type YahooFinanceResponse struct {
	Chart Chart
}

type Chart struct {
	Result []Result
}

type Result struct {
	Indicators Indicators
	Timestamp  []int64
}

type Indicators struct {
	Quote []Quote
}

type Quote struct {
	Close  []float64
	Open   []float64
	High   []float64
	Low    []float64
	Volume []int64
}
