package main

import (
	"log"
)

var (
	bc = NewBlockchain()
)

func main() {
	// init
	bc.AddBlock("Ali send 1 Coin to Nasro")
	bc.AddBlock("Manel send 1 Coin to Ahmed")

	for _, b := range bc.Blocks {
		b.ShowBlock()
	}

	if err := RunServer(); err != nil {
		log.Fatal("Can't run the server")
	}
}
