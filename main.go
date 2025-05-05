package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const subsidy = 10

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Data          []byte
	PrevHash      []byte
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

type CLI struct {
	bc *BlockChain
}

type Transaction struct {
	ID   []byte
	Vin  []TxtInput
	Vout []TXOutput
}

type TXOutput struct {
	Value        int
	ScriptPubKey string
}

type TxtInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  []*Transaction{},
		PrevBlockHash: prevHash,
		Hash:          []byte{},
		Data:          []byte(data),
		PrevHash:      prevHash,
	}
	block.Hash = []byte(fmt.Sprintf("%x", data+string(prevHash)))
	return block
}

func Genesis() *Block {
	return NewBlock("Genesis", []byte{})
}

func (b *Block) Serialize() []byte {
	return append(b.PrevHash, b.Data...)
}

func DeserializeBlock(data []byte) *Block {
	return &Block{
		Data:     data[32:],
		PrevHash: data[:32],
		Hash:     []byte{},
	}
}

func (bc *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			return err
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}

func NewBlockchain() *BlockChain {
	var tip []byte

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			genesis := Genesis()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Fatal(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Fatal(err)
			}
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
			return err
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return &BlockChain{tip, db}
}

func (bc *BlockChain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.tip, bc.db}
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	i.currentHash = block.PrevHash
	return block
}

func (cli *CLI) printChain() {
	iter := cli.bc.Iterator()

	for {
		block := iter.Next()
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Block added.")
}

func (cli *CLI) Run() {
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	if len(os.Args) < 2 {
		fmt.Println("expected 'addblock' or 'printchain' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "addblock":
		addBlockCmd.Parse(os.Args[2:])
	case "printchain":
		printChainCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func main() {
	cli := CLI{NewBlockchain()}
	cli.Run()
}

func NewCoinBaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TxtInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TxtInput{txin}, []TXOutput{txout}}
	return &tx
}
