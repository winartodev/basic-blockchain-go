package main

import (
	"fmt"
	"strconv"

	"github.com/winartodev/basic-blockchain-go/core"
)

func main() {
	// Initialize Blockchain
	bc := core.NewBlockchain()

	// Adding new blocks
	bc.AddBlock("Block 1")
	bc.AddBlock("Block 2")

	// Show all blocks in the blockchain
	for _, block := range bc.Blocks {
		fmt.Printf("PrevHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := core.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
