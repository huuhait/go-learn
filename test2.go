package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect("nats://127.0.0.1:4222")

	nc.Subscribe("test-gui-di", func(msg *nats.Msg) {
		log.Println(msg.Data)
	})

	for {
		time.Sleep(10 * time.Second)
	}
}
