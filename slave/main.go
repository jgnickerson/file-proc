package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

var nc *nats.Conn
var err error

func main() {
	nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := nc.Request("file.server", nil, time.Second * 30)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(resp.Data))
}
