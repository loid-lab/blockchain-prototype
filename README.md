# 🧱 Simple Blockchain with Proof of Work and CLI in Go

This is a minimal blockchain implementation written in Go. It demonstrates how a basic blockchain works by chaining blocks together using cryptographic hashes and includes a simple Proof-of-Work (PoW) mechanism for mining new blocks. It uses **BoltDB** for persistent storage and provides a basic **command-line interface (CLI)** for interacting with the blockchain.

---

## 🔧 Features

- ✅ Creates a genesis block
- ✅ Adds new blocks to a persistent chain
- ✅ CLI to interact with the chain
- ✅ Each block includes:
  - Data (e.g. transactions)
  - Hash of the previous block
  - Its own SHA-256 hash (simplified)
  - Nonce (for PoW)
- ✅ Implements Proof-of-Work:
  - Adjustable difficulty via `targetBits`
  - Mines blocks by finding a valid nonce
  - Validates PoW on each block
- ✅ Uses BoltDB for local data storage
- ✅ Added support for transactions and transaction rewards
- ✅ Added wallet functionality with public/private key pairs
- ✅ Network Support:
  - Starts a TCP server node
  - Nodes can send and receive blocks and transactions
  - Handles new block propagation across the network
  - Includes basic peer-to-peer functionality

---

## 🧠 How It Works

Each block contains data and references the previous block's hash. A SHA-256 hash is computed from the block's data, previous hash, and a nonce. The Proof-of-Work algorithm ensures the block hash is below a target value, enforcing difficulty and securing the blockchain. 

The blockchain now supports transactions, where each block may contain one or more transactions. Each wallet generates public/private keys and is used to manage transactions and block mining. A peer-to-peer network allows nodes to communicate, propagate new blocks and transactions, and maintain consensus.

---

## 📦 Getting Started

### ✅ Prerequisites

Make sure you have Go installed:  
👉 https://golang.org/dl/

---

### 🏃 Run the Program

Clone the repository:

```bash
git clone https://github.com/loid-lab/blockchain.git
cd blockchain
go run main.go
```

---

### 🧪 CLI Usage

| Command             | Description                        |
|---------------------|------------------------------------|
| `addblock -data=X`  | Adds a block with transaction data |
| `createwallet`      | Creates a new wallet (public/private key pair) |
| `miningreward`      | Creates a mining reward transaction (coinbase) |
| `printchain`        | Prints all blocks in the blockchain |

#### Example

```bash
go run main.go addblock -data="Send 1 BTC to Ivan"
go run main.go addblock -data="Send 2 more BTC to Ivan"
go run main.go printchain
```

#### Sample Output

```
Prev. hash: 
Data: Genesis
Hash: abc123...

Prev. hash: abc123...
Data: Send 1 BTC to Ivan
Hash: def456...

Prev. hash: def456...
Data: Send 2 more BTC to Ivan
Hash: ghi789...
```

---

## 📝 Notes

This is a **learning** and **demo** project. It does not include:
- Real cryptographic hash functions or difficulty adjustments
- Full consensus or network security (basic P2P is included)
- Wallets or real transactions
- Security, encryption, or signatures

---

## 📚 Resources

- 📘 [Mastering Bitcoin – Andreas M. Antonopoulos](https://github.com/bitcoinbook/bitcoinbook)
- 📘 [Go Programming Language](https://golang.org/doc/)
- 📘 [BoltDB](https://github.com/boltdb/bolt)

---

## 📄 License

This project is licensed under the **MIT License**.  
Feel free to fork and improve!