
package main

import (
	"fmt"
	"github.com/Areebxh/assignment01bca"
)

func main() {
	blockchain := &assignment01bca.Blockchain{}

	blockchain.Blocks = append(blockchain.Blocks, assignment01bca.NewBlock("Alice to Bob", 123, "genesis"))
	blockchain.Blocks = append(blockchain.Blocks, assignment01bca.NewBlock("Bob to Charlie", 456, blockchain.Blocks[len(blockchain.Blocks)-1].CurrentHash))

	
	fmt.Println("Before Change:")
	blockchain.DisplayBlocks()

	
	blockchain.ChangeBlock(1, "Mallory to Eve")	

	fmt.Println("\nAfter Change:")
	blockchain.DisplayBlocks()

	isValid := blockchain.VerifyChain()
	fmt.Println("\nIs Blockchain Valid?", isValid)
}
