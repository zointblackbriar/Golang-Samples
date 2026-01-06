package main

import (
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	bc := NewBlockchain("address")

	if len(bc.blocks) != 1 {
		t.Errorf("Expected 1 block, got %d", len(bc.blocks))
	}

	genesisBlock := bc.blocks[0]
	if genesisBlock.Height != 0 {
		t.Errorf("Expected genesis block height to be 0, got %d", genesisBlock.Height)
	}

	if len(genesisBlock.Transactions) != 1 {
		t.Errorf("Expected 1 transaction in genesis block, got %d", len(genesisBlock.Transactions))
	}

	coinbaseTx := genesisBlock.Transactions[0]
	if !coinbaseTx.IsCoinbase() {
		t.Error("Expected the transaction in genesis block to be a coinbase transaction")
	}
}

func TestAddBlock(t *testing.T) {
	bc := NewBlockchain("address")
	tx := NewCoinbaseTX("address2", "data")
	transactions := []*Transaction{tx}
	bc.AddBlock(transactions)

	if len(bc.blocks) != 2 {
		t.Errorf("Expected 2 blocks, got %d", len(bc.blocks))
	}

	newBlock := bc.blocks[1]
	if newBlock.Height != 1 {
		t.Errorf("Expected new block height to be 1, got %d", newBlock.Height)
	}

	if len(newBlock.Transactions) != 1 {
		t.Errorf("Expected 1 transaction in new block, got %d", len(newBlock.Transactions))
	}

	if string(newBlock.Transactions[0].ID) != string(tx.ID) {
		t.Error("Transaction in new block does not match")
	}
}

func TestGetBalance(t *testing.T) {
	bc := NewBlockchain("address")
	balance := bc.GetBalance("address")
	if balance != 10 {
		t.Errorf("Expected balance to be 10, got %d", balance)
	}
}

func TestFindUnspentTransactions(t *testing.T) {
	bc := NewBlockchain("address")
	unspentTXs := bc.FindUnspentTransactions("address")

	if len(unspentTXs) != 1 {
		t.Errorf("Expected 1 unspent transaction, got %d", len(unspentTXs))
	}
}

func TestFindUTXO(t *testing.T) {
	bc := NewBlockchain("address")
	utxos := bc.FindUTXO("address")

	if len(utxos) != 1 {
		t.Errorf("Expected 1 UTXO, got %d", len(utxos))
	}
}

func TestMineBlock(t *testing.T) {
	bc := NewBlockchain("address")
	tx := NewCoinbaseTX("address2", "")
	transactions := []*Transaction{tx}
	bc.MineBlock(transactions)

	if len(bc.blocks) != 2 {
		t.Errorf("Expected 2 blocks after mining, got %d", len(bc.blocks))
	}
}

func TestFindTransaction(t *testing.T) {
	bc := NewBlockchain("address")
	tx := bc.blocks[0].Transactions[0]

	foundTx, err := bc.FindTransaction(tx.ID)
	if err != nil {
		t.Errorf("Could not find transaction: %v", err)
	}

	if string(foundTx.ID) != string(tx.ID) {
		t.Error("Found transaction does not match original transaction")
	}
}

func TestGetBlocks(t *testing.T) {
	bc := NewBlockchain("address")
	blocks := bc.GetBlocks()

	if len(blocks) != 1 {
		t.Errorf("Expected 1 block, got %d", len(blocks))
	}
}

func TestVerifyTransaction(t *testing.T) {
	bc := NewBlockchain("address")
	tx := NewCoinbaseTX("address", "")

	if !bc.VerifyTransaction(tx) {
		t.Error("Coinbase transaction should be valid")
	}
}

func TestFindSpendableOutputs(t *testing.T) {
	bc := NewBlockchain("address")
	accumulated, unspentOutputs := bc.FindSpendableOutputs("address", 10)

	if accumulated < 10 {
		t.Errorf("Expected accumulated amount to be at least 10, got %d", accumulated)
	}

	if len(unspentOutputs) == 0 {
		t.Error("Expected to find spendable outputs")
	}
}