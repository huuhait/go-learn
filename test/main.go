package main

import (
	"log"

	"github.com/shopspring/decimal"
)

type WithdrawLimit struct {
	ID       int64  `json:"id"`
	Group    string `json:"group"`
	KycLevel string `json:"kyc_level"`
}

func main() {
	a := decimal.NewFromInt(3)
	b := decimal.NewFromFloat(4.5)

	// c := a.Add(b)

	if a.Cmp(b) == 0 {
		log.Println("a == b")
	} else if a.Cmp(b) == 1 {
		log.Println("a > b")
	} else if a.Cmp(b) == -1 {
		log.Println("a < b")
	}

	// log.Println(c)
}
