
package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)


type Block struct {
	Transaction   string
	Nonce         int
	PreviousHash  string
	CurrentHash   string
}


type Blockchain struct {
	Blocks []*Block
}


func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		CurrentHash:  CalculateHash(transaction, nonce, previousHash),
	}

	return block
}


func (bc *Blockchain) DisplayBlocks() {
	for _, block := range bc.Blocks {
		fmt.Printf("Transaction: %s\nNonce: %d\nPrevious Hash: %s\nCurrent Hash: %s\n\n",
			block.Transaction, block.Nonce, block.PreviousHash, block.CurrentHash)
	}
}


func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex >= 0 && blockIndex < len(bc.Blocks) {
		bc.Blocks[blockIndex].Transaction = newTransaction
		bc.Blocks[blockIndex].CurrentHash = CalculateHash(bc.Blocks[blockIndex].Transaction, bc.Blocks[blockIndex].Nonce, bc.Blocks[blockIndex].PreviousHash)
	}
}


func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		if bc.Blocks[i].PreviousHash != bc.Blocks[i-1].CurrentHash {
			return false
		}
	}
	return true
}


func CalculateHash(transaction string, nonce int, previousHash string) string {
	data := fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)
	hashInBytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashInBytes[:])
}
