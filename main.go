package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Timestamp     int64  // Timestamp when the block was created
	Data          []byte // Actual data stored in the block (e.g., transactions)
	PrevBlockHash []byte // Hash of the previous block in the chain
	Hash          []byte // Hash of the current block (calculated from its contents)
}

// BlockChain represents the chain of blocks
type BlockChain struct {
	blocks []*Block // Slice of pointers to blocks
}

// SetHash calculates and sets the hash of a block using SHA-256
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10)) // Convert timestamp to byte slice
	// Concatenate previous hash, data, and timestamp
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers) // Generate SHA-256 hash

	b.Hash = hash[:] // Set the block's hash
}

// NewBlock creates a new block with the given data and previous block's hash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(), // Current time in Unix format
		Data:          []byte(data),      // Convert data to bytes
		PrevBlockHash: prevBlockHash,     // Link to previous block
		Hash:          []byte{},          // Will be set below
	}
	block.SetHash() // Calculate the block’s hash
	return block
}

// AddBlock adds a new block with given data to the blockchain
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]   // Get the latest block
	newBlock := NewBlock(data, prevBlock.Hash) // Create a new block linked to it
	bc.blocks = append(bc.blocks, newBlock)    // Append the new block to the chain
}

// NewGenesisBlock creates the first block in the blockchain (with no predecessor)
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) // No previous hash for the genesis block
}

// NewBlockchain initializes a new blockchain with the genesis block
func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

// main is the entry point of the program
func main() {
	bc := NewBlockchain() // Initialize blockchain

	// Add two more blocks
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	// Print out each block's details
	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
