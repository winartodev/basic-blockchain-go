package core

// Blockchain is a data structure that stores a sequence of blocks.
type Blockchain struct {
	Blocks []*Block // a collection of block data structures
}

// AddBlock adds a new block to the blockchain
// while maintaining the reference to the previous block.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock("GG Blockchain", []byte{})}}
}
