package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"strconv"
	"time"
)

// Block is used to store valuable information in a blockchain.
type Block struct {
	Timestamp     int64  // Store current timestamp when block is created
	Data          []byte // Contains the actual valuable information within the block.
	PrevBlockHash []byte // Stores the hash of the previous block.
	Hash          []byte // Hash of the current block.
	Nonce         int
}

// SetHash calculates the SHA-256 hash based on the concatenated header information.
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock creates a new block while preserving the reference to the previous block's hash.
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         0,
	}
	pow := NewProofOfWork(block)
	nonce, hash, _ := pow.Mine()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGenesisBlock(data string, hash []byte) *Block {
	return NewBlock(data, hash)
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		return nil
	}

	return result.Bytes()
}
