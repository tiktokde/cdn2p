package main

import (
	"log"

	cdn2proxy "github.com/tiktokde/cdn2p"
)

func main() {
	err := cdn2proxy.StartProxy("127.0.0.1:10888", "ws://127.0.0.1:9000/ws", "", "https://9.9.9.9/dns-query")
	if err != nil {
		log.Fatal(err)
	}
}
