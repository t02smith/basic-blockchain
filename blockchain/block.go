package blockchain

// methods

// static

// generates a new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
		Nonce:    0,
	}

	pow := CreateProofOfWork(block)
	nonce, hash := RunProofOfWork(pow)

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
