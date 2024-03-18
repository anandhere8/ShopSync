package main

import (
	server "github.com/anandhere8/ShopSync/cmd/shopsync"
)

const (
	address = "0.0.0.0:8080"
)

func main() {

	server := server.NewServer()
	server.Start(address)
}
