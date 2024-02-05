package main

import (
	"encoding/base64"
	"fmt"
)

func (bc *Blockchain) printChain() {
	// Implement the logic to print all the blocks of the blockchain

	if len(bc.blocks) == 0 {
		fmt.Println("No blocks in the blockchain")
		return
	}

	
	for _, block := range bc.blocks {
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: \n")
		
		for i := 0; i < len(block.Transactions); i++ {
			fmt.Printf("- Transaction: %d\n", i+1)
			fmt.Printf("  	+ Receiver: %s\n", block.Transactions[i].Receiver)
			fmt.Printf("  	+ Sender: %s\n", block.Transactions[i].Sender)
			encodingStr := base64.StdEncoding.EncodeToString(block.Transactions[i].Signature)
			fmt.Printf("  	+ Signature: %s\n", encodingStr)
		}
		fmt.Printf("HashRoot: %x\n", block.HashRoot)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}