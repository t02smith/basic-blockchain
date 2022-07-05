package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// methods

// Generates the hash for a block
func (b *Block) HashBlock() {
	hash := sha256.Sum256(bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{}))
	b.Hash = hash[:]
}

// static

// generates a new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
	}

	block.HashBlock()
	return block
}
