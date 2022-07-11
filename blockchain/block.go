package blockchain

import (
	"bytes"
	"encoding/gob"
)

// methods

// static

// generates a new block
func CreateBlock(txns []*Transaction, prevHash []byte) *Block {
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

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var b Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&b)
	Handle(err)

	return &b
}
