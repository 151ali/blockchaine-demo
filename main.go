package main

import (
	"log"
)

var (
	bc = NewBlockchain()
)

func main() {
	// init
	bc.AddBlock(Transaction{"Ali", "Nasro", 1})
	bc.AddBlock(Transaction{"Ahmed", "Manel", 3})

	for _, b := range bc.Blocks {
		b.ShowBlock()
	}

	if err := RunServer(); err != nil {
		log.Fatal("Can't run the server")
	}
}
