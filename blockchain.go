package main

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]

	// Create a new transaction using the data
	transaction := &Transaction{Data: []byte(data)}

	// Pass the transaction and previous block's hash to NewBlock
	newBlock := NewBlock([]*Transaction{transaction}, prevBlock.Hash)

	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
