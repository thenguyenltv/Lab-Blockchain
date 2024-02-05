package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64          // Current timestamp (when the block is created)
	Transactions  []*Transaction // Actual valuable information containing in the block
	PrevBlockHash []byte         // Hash of the previous block
	HashRoot	  []byte         // Hash of the merkle tree root
	Hash          []byte         // Hash of this block
}

type Transaction struct {
	Sender   		[]byte // Sender's address
	Receiver 		[]byte // Receiver's address
	Signature     	[]byte // Signature of the transaction
}

func (b *Block) SetHash() {
	str := strconv.FormatInt(b.Timestamp, 10)
	timestampBytes := []byte(str)
	data := HashTransaction(b.Transactions)

	headers := bytes.Join([][]byte{timestampBytes, data, b.PrevBlockHash, b.HashRoot}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func HashTransaction(Transactions []*Transaction) []byte {
	var concatenatedHash []byte
	for _, tx := range Transactions {
		hash := sha256.Sum256(bytes.Join(
			[][]byte{
				tx.Receiver, 
				tx.Sender, 
				tx.Signature,
			},
			[]byte{}),
		) // Hash each transaction

		// concatenatedHash = bytes.Join([][]byte{concatenatedHash, hash[:]}, []byte{})
		concatenatedHash = append(concatenatedHash, hash[:]...) // Concatenate all transaction hashes
	}

	hash := sha256.Sum256(concatenatedHash) // Hash the concatenated hashes
	return hash[:]
}

func NewBlock(Transactions []*Transaction, prevBlockHash, HashRoot []byte) *Block {
	block := &Block{Timestamp: time.Now().Unix(), Transactions: Transactions, PrevBlockHash: prevBlockHash, HashRoot: HashRoot, Hash: []byte{}}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	var twoDSlice [][]byte

	//initialize the genesis block with a string, 
	dataStr := []byte("Genesis Block")
	twoDSlice = append(twoDSlice, dataStr)
	transactions := []*Transaction{&Transaction{nil, nil, dataStr}}
	
	return NewBlock(transactions, []byte{}, NewMerkleTree(twoDSlice).NodeRoot.Data)
}

