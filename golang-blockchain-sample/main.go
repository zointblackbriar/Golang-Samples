package main

import (
	"fmt"
	"strconv"
)

const genesisData = "Genesis Block"

func main() {
	// Create addresses
	address1 := "Alice"
	address2 := "Bob"
	address3 := "Charlie"

	fmt.Println("=== Creating Blockchain ===")
	bc := NewBlockchain(address1)
	defer fmt.Println("\n=== Blockchain Demo Complete ===")

	fmt.Printf("\nGenesis block created. %s receives mining reward of %d\n", address1, subsidy)

	// Check initial balance
	fmt.Printf("\nBalance of %s: %d\n", address1, bc.GetBalance(address1))

	// Create and mine first transaction
	fmt.Println("\n=== Transaction 1: Alice sends 3 to Bob ===")
	tx1 := NewUTXOTransaction(address1, address2, 3, bc)
	bc.MineBlock([]*Transaction{tx1})
	fmt.Printf("%s mining reward: %d\n", address1, subsidy)

	// Check balances
	fmt.Printf("\nBalance of %s: %d\n", address1, bc.GetBalance(address1))
	fmt.Printf("Balance of %s: %d\n", address2, bc.GetBalance(address2))

	// Create and mine second transaction
	fmt.Println("\n=== Transaction 2: Alice sends 2 to Charlie ===")
	tx2 := NewUTXOTransaction(address1, address3, 2, bc)
	bc.MineBlock([]*Transaction{tx2})
	fmt.Printf("%s mining reward: %d\n", address1, subsidy)

	// Check balances
	fmt.Printf("\nBalance of %s: %d\n", address1, bc.GetBalance(address1))
	fmt.Printf("Balance of %s: %d\n", address2, bc.GetBalance(address2))
	fmt.Printf("Balance of %s: %d\n", address3, bc.GetBalance(address3))

	// Create and mine third transaction
	fmt.Println("\n=== Transaction 3: Bob sends 1 to Charlie ===")
	tx3 := NewUTXOTransaction(address2, address3, 1, bc)
	bc.MineBlock([]*Transaction{tx3})
	fmt.Printf("%s mining reward: %d\n", address2, subsidy)

	// Final balances
	fmt.Println("\n=== Final Balances ===")
	fmt.Printf("Balance of %s: %d\n", address1, bc.GetBalance(address1))
	fmt.Printf("Balance of %s: %d\n", address2, bc.GetBalance(address2))
	fmt.Printf("Balance of %s: %d\n", address3, bc.GetBalance(address3))

	// Print blockchain details
	fmt.Println("\n=== Blockchain Details ===")
	blocks := bc.GetBlocks()
	for i, block := range blocks {
		fmt.Printf("\n--- Block %d ---\n", i)
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Transactions: %d\n", len(block.Transactions))

		// Validate proof of work
		pow := NewProofOfWork(block)
		fmt.Printf("PoW Valid: %s\n", strconv.FormatBool(pow.Validate()))

		// Print transaction details
		for j, tx := range block.Transactions {
			fmt.Printf("\nTransaction %d:\n", j)
			fmt.Print(tx.String())
		}
	}

	// Verify blockchain integrity
	fmt.Println("\n=== Verifying Blockchain Integrity ===")
	isValid := true
	for i := 1; i < len(blocks); i++ {
		currentBlock := blocks[i]
		previousBlock := blocks[i-1]

		// Check if previous hash matches
		if string(currentBlock.PreviousHash) != string(previousBlock.Hash) {
			fmt.Printf("Block %d: Previous hash mismatch!\n", i)
			isValid = false
		}

		// Validate proof of work
		pow := NewProofOfWork(currentBlock)
		if !pow.Validate() {
			fmt.Printf("Block %d: Invalid proof of work!\n", i)
			isValid = false
		}
	}

	if isValid {
		fmt.Println("✓ Blockchain is valid!")
	} else {
		fmt.Println("✗ Blockchain is invalid!")
	}
}