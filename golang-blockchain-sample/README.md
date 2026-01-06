# Golang Blockchain Implementation

A complete blockchain implementation in Go featuring proof-of-work consensus, transaction handling, and UTXO (Unspent Transaction Output) model.

## Features

- **Proof of Work (PoW)**: Mining algorithm with adjustable difficulty
- **Transaction System**: Full UTXO model with inputs and outputs
- **Blockchain Structure**: Linked blocks with cryptographic hashing
- **Coinbase Transactions**: Mining rewards for block creation
- **Balance Tracking**: Query balances for any address
- **Transaction Verification**: Validate transactions before adding to blockchain
- **Chain Integrity Validation**: Verify the entire blockchain

## Architecture

### Components

1. **Block** (`block.go`)
   - Contains transactions, timestamp, previous hash, and nonce
   - Implements proof-of-work mining
   - Uses SHA-256 hashing

2. **Blockchain** (`blockchain.go`)
   - Manages the chain of blocks
   - Handles UTXO tracking
   - Validates and mines new blocks
   - Provides balance queries

3. **Transaction** (`transaction.go`)
   - UTXO-based transaction model
   - Supports coinbase (mining reward) transactions
   - Transaction inputs and outputs
   - Digital signature verification (simplified)

## How It Works

### Proof of Work
The blockchain uses a proof-of-work algorithm where miners must find a nonce that produces a hash with a specific number of leading zeros (difficulty). The difficulty is set by `targetBits = 16`.

### UTXO Model
Transactions use the Unspent Transaction Output (UTXO) model:
- **Inputs**: Reference previous transaction outputs
- **Outputs**: Create new spendable amounts
- **Change**: Excess amount returned to sender

### Mining Rewards
Block miners receive a subsidy of 10 coins for each mined block.

## Usage

### Build
```bash
go build -o blockchain
```

### Run
```bash
./blockchain
```

Or directly:
```bash
go run .
```

## Example Output

The demo creates a blockchain with multiple transactions:

1. **Genesis Block**: Creates initial coins for Alice
2. **Transaction 1**: Alice sends 3 coins to Bob
3. **Transaction 2**: Alice sends 2 coins to Charlie
4. **Transaction 3**: Bob sends 1 coin to Charlie

The program displays:
- Mining progress with hash attempts
- Transaction details
- Balance updates
- Complete blockchain structure
- Blockchain validation results

## Code Structure

```
.
├── block.go          # Block structure and proof-of-work
├── blockchain.go     # Blockchain management
├── transaction.go    # Transaction handling
├── main.go          # Demo application
├── go.mod           # Go module definition
└── README.md        # This file
```

## Key Concepts

### Block Structure
```go
type Block struct {
    Timestamp     int64
    Transactions  []*Transaction
    PreviousHash  []byte
    Hash          []byte
    Nonce         int
    Height        int
}
```

### Transaction Structure
```go
type Transaction struct {
    ID   []byte
    Vin  []TXInput   // Inputs
    Vout []TXOutput  // Outputs
}
```

## Technical Details

- **Hashing Algorithm**: SHA-256
- **Difficulty**: 16 bits (adjustable via `targetBits`)
- **Subsidy**: 10 coins per block
- **Encoding**: Gob encoding for transaction serialization

## Security Features

1. **Cryptographic Hashing**: Each block is linked via SHA-256 hashes
2. **Proof of Work**: Prevents easy block creation
3. **Transaction Validation**: Ensures sufficient funds before transactions
4. **Chain Validation**: Verifies entire blockchain integrity
5. **UTXO Tracking**: Prevents double-spending

## Extending the Blockchain

To add more features, consider:
- Persistent storage (database)
- Network layer (P2P)
- Advanced cryptography (digital signatures with ECDSA)
- Merkle trees for transaction verification
- Smart contract support
- API endpoints (REST/gRPC)

## Requirements

- Go 1.13 or higher
- No external dependencies for core functionality

## License

MIT License

---

**Note**: This is an educational implementation. Production blockchains require additional security measures, networking capabilities, and optimizations.
