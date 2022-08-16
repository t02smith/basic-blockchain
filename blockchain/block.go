package blockchain

import (
	"bytes"
	"encoding/gob"
)

// generates a new block
func createBlock(txns []*Transaction, prevHash []byte) *Block {
	block := &Block{
		Hash:         []byte{},
		Transactions: txns,
		PrevHash:     prevHash,
		Nonce:        0,
	}

	pow := CreateProofOfWork(block)
	nonce, hash := RunProofOfWork(pow)

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// IO

func (b *Block) serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	handle(err)

	return res.Bytes()
}

func deserialize(data []byte) *Block {
	var b Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&b)
	handle(err)

	return &b
}
