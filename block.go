package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Transaction struct {
	Data []byte
}

// Block keeps block headers
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction 
	PrevBlockHash []byte
	Hash          []byte
}

// SetHash calculates and sets block hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	var data [][]byte

	for _, transaction := range b.Transactions {
		data = append(data, transaction.Data)
	}
	
	headers := bytes.Join([][]byte{b.PrevBlockHash, bytes.Join(data, []byte{}), timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock creates and returns Block
func NewBlock(transaction []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transaction, prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	// Create a coinbase transaction (you need to provide a valid coinbase transaction here)
	coinbase := &Transaction{Data: []byte("Genesis Transaction")}
	return NewBlock([]*Transaction{coinbase}, []byte{})
}