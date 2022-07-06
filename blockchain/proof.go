package blockchain

import (
	"bytes"
	"crypto/sha256"
	"log"
	"math"
	"math/big"
)

const DIFFICULTY = 15

// methods

// static

// runs the proof of work function to find the nonce + hash
func RunProofOfWork(pow *ProofOfWork) (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 {
		data := CreateNonce(pow, nonce)
		hash = sha256.Sum256(data)

		log.Printf("%x\n", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	log.Println()
	return nonce, hash[:]
}

// validates that a pow is valid
func ValidateProofOfWork(pow *ProofOfWork) bool {
	var intHash big.Int

	data := CreateNonce(pow, pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}

// creates a nonce given a ProofOfWork
func CreateNonce(pow *ProofOfWork, nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(DIFFICULTY)),
		},
		[]byte{},
	)
	return data
}

// generates a new proof of work
func CreateProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-DIFFICULTY))

	pow := &ProofOfWork{
		Block:  b,
		Target: target,
	}
	return pow
}
