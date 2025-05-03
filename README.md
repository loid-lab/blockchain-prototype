# Simple Blockchain with Proof of Work in Go

This is a minimal blockchain implementation written in Go. It demonstrates how a basic blockchain works by chaining blocks together using cryptographic hashes and includes a simple Proof-of-Work (PoW) mechanism for mining new blocks.

## 🔧 Features

- Creates a genesis block
- Adds new blocks to the chain
- Each block includes:
  - Timestamp
  - Data (e.g. transactions)
  - Hash of the previous block
  - Its own SHA-256 hash
  - Nonce (for PoW)
- Implements Proof-of-Work:
  - Adjusts difficulty using `targetBits`
  - Mines blocks by finding a valid nonce
  - Validates the PoW of each block

## 🧱 How It Works

Each block stores data and references the hash of the previous block. A SHA-256 hash is computed from the block content, the previous block's hash, and a nonce. The PoW algorithm ensures the resulting hash is below a target value, which enforces difficulty and secures the chain.

## 📦 Getting Started

### Prerequisites

Make sure you have Go installed: https://golang.org/dl/

### Run the Program

Clone the repository and run:

```bash
go run main.go
### Sample Output

```
Prev. hash: 
Data: Genesis Block
Hash: e3f1f3...

Prev. hash: e3f1f3...
Data: Send 1 BTC to Ivan
Hash: a1b2c3...

Prev. hash: a1b2c3...
Data: Send 2 more BTC to Ivan
Hash: d4e5f6...
```

## 📝 Notes

This is a very simplified example and is not suitable for production use. It doesn't include:
- Networking
- Transactions
- Security layers

## 📚 Learn More

- [Mastering Bitcoin by Andreas Antonopoulos](https://github.com/bitcoinbook/bitcoinbook)
- [Go Documentation](https://golang.org/doc/)

## 📄 License

This project is licensed under the MIT License.
