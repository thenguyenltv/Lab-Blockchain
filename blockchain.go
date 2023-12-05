package main

import (
	"bytes"
	"errors"
)

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	blocks []*Block
}


// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.blocks[len(bc.blocks)-1]

	bc.blocks = append(bc.blocks, NewBlock(transactions, prevBlock.Hash, prevBlock.HashRoot))

	// create new merkle tree
	var blocks [][]byte
	for _, block := range bc.blocks {
		blocks = append(blocks, block.Hash)
	}

	HashRoot := NewMerkleTree(blocks).NodeRoot.Data
	_ = HashRoot // Ignore the unused variable error

	bc.blocks[len(bc.blocks)-1].HashRoot = HashRoot
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// CurrentBlock returns the last block
func (bc Blockchain) CurrentBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

// GetBlock returns the block of a given hash
func (bc Blockchain) GetBlock(hash []byte) (*Block, error) {
	for i := len(bc.blocks) - 1; i >= 0; i-- {
		if bytes.Equal(bc.blocks[i].Hash, hash) {
			return bc.blocks[i], nil
		}
	}

	return nil, errors.New("no blocks has the given hash")
}
