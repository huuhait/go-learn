package jobs

import (
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/huuhait/go-learn/config"
)

type FetchPrice struct{}

func NewFetchPrice() *FetchPrice {
	return &FetchPrice{}
}

func (FetchPrice) Process() {
	http_client := resty.New()

	res, err := http_client.R().Get("https://min-api.cryptocompare.com/data/pricemulti?fsyms=USD,USDT&tsyms=USD,USDT,EUR,VND,CNY,JPY")
	if err != nil {
		log.Println("Failed to fetch price from global")
		return
	}

	config.RedisClient.Set("global_price", string(res.Body()))

	log.Println("DONE")

	time.Sleep(20 * time.Second)
}
