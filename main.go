package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	var data [][]byte

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		for _, transaction := range block.Transactions {
			data = append(data, transaction.Data)
		}
		fmt.Printf("Transactions: %s\n", data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
