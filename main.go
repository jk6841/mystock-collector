package main

import (
	"fmt"
	"mystock-collector/mystock"
	"time"
)

func main() {
	stock := mystock.Stock{
		Symbol: "AAPL",
		Name:   "Apple",
	}
	for {
		now := time.Now()
		data := stock.GetData(now.Add(time.Minute*(-1)), now)
		fmt.Println(data)
		time.Sleep(time.Minute)
	}
}
