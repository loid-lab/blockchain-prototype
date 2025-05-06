// Optimized blockchain code with retained functionality
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"math/big"
	"time"
)

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
}

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

type Blockchain struct {
	blocks []*Block
}

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	err := gob.NewEncoder(&result).Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block
	err := gob.NewDecoder(bytes.NewReader(d)).Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(transactions, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlock(transactions []*Transaction, prevHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevHash,
	}
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock([]*Transaction{{"genesis", "satoshi", 50}}, []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, 256-16) // Difficulty
	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.hashTransactions(),
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(nonce)),
	}, []byte{})
	return data
}

func (pow *ProofOfWork) hashTransactions() []byte {
	var txHashes [][]byte
	for _, tx := range pow.block.Transactions {
		txData := []byte(fmt.Sprintf("%s%s%f", tx.Sender, tx.Recipient, tx.Amount))
		hash := sha256.Sum256(txData)
		txHashes = append(txHashes, hash[:])
	}
	hash := sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return hash[:]
}

func (pow *ProofOfWork) Run() ([]byte, int) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	for {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		}
		nonce++
	}
	return hash[:], nonce
}

func IntToHex(n int64) []byte {
	return []byte(fmt.Sprintf("%x", n))
}

func main() {
	bc := NewBlockchain()
	tx1 := &Transaction{"Alice", "Bob", 10}
	tx2 := &Transaction{"Bob", "Charlie", 5}
	bc.AddBlock([]*Transaction{tx1, tx2})
	for _, block := range bc.blocks {
		fmt.Printf("\nTimestamp: %d\n", block.Timestamp)
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		for _, tx := range block.Transactions {
			fmt.Printf("Transaction: %s -> %s: %.2f\n", tx.Sender, tx.Recipient, tx.Amount)
		}
	}
}
