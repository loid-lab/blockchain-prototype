# Simple Blockchain Prototype in Go

This is a minimal blockchain implementation written in Go. It demonstrates how a basic blockchain works by chaining blocks together using cryptographic hashes.

## 🔧 Features

- Creates a genesis block
- Adds new blocks to the chain
- Each block contains:
  - Timestamp
  - Data (e.g. transactions)
  - Hash of the previous block
  - Its own SHA-256 hash

## 🧱 How It Works

Each block stores data and references the hash of the previous block. When a block is added, a hash is computed from its content and the previous block’s hash, ensuring immutability.

## 📦 Getting Started

### Prerequisites

Make sure you have Go installed: https://golang.org/dl/

### Run the Program

Clone the repository and run:

```bash
go run main.go
```

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
- Proof of Work
- Networking
- Transactions
- Security layers

## 📚 Learn More

- [Mastering Bitcoin by Andreas Antonopoulos](https://github.com/bitcoinbook/bitcoinbook)
- [Go Documentation](https://golang.org/doc/)

## 📄 License

This project is licensed under the MIT License.
