package mystock

type Candle struct {
	Stock  Stock
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume int64
	Start  int64
	End    int64
}
