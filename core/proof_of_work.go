package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/winartodev/basic-blockchain-go/utils"
)

var maxNonce = math.MaxInt64

const (
	// TargetBits defines the number of leading zero bits required in the hash
	TargetBits = 24
)

// ProofOfWork holds a pointer to a Block and a pointer to Target
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TargetBits))

	return &ProofOfWork{
		Block:  b,
		Target: target,
	}
}

// PrepareData used to merge block fields with the target and nonce
// 	`nonce` is counter from the Hashcash
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			utils.IntegerToHex(pow.Block.Timestamp),
			utils.IntegerToHex(int64(TargetBits)),
			utils.IntegerToHex(int64(nonce)),
		}, []byte{},
	)

	return data
}

// Mine is copre of the PoW algorithm
func (pow *ProofOfWork) Mine() (int, []byte, error) {
	var hashInt big.Int
	var hash [32]byte

	fmt.Printf("Mining the block containing \"%s\"\n", pow.Block.Data)

	startTime := time.Now()
	defer func() {
		duration := time.Since(startTime)
		fmt.Printf("\nMining took %v\n\n", duration)
	}()

	for nonce := 0; nonce < maxNonce; nonce++ {
		data := pow.PrepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.Target) == -1 {
			return nonce, hash[:], nil
		}
	}

	fmt.Print("\n\n")

	return 0, nil, fmt.Errorf("mining exhausted")
}

// Validate used to validating block with proof of works
func (pow *ProofOfWork) Validate() bool {
	data := pow.PrepareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)

	var hashInt big.Int
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.Target) == -1
}
