package main

import "crypto/sha256"

type MerkleTree struct {
	NodeRoot *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// NewMerkleNode creates a new Merkle tree node
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}

	node.Left = left
	node.Right = right

	return &node
}

// NewMerkleTree creates a new Merkle tree from a sequence of data
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	// If there is an odd number of data items, duplicate the last one
	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	// Create a leaf node for each data item
	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	// Create the parent nodes by combining the leaf nodes
	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode

		// Combine each pair of nodes
		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}

		// Set the new level as the current level
		nodes = newLevel
	}

	mTree := MerkleTree{&nodes[0]}

	return &mTree
}

