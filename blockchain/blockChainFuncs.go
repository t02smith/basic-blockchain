package blockchain

// methods

// returns the most recent node in the blockchain
func (chain *BlockChain) Head() *Block {
	return chain.Blocks[len(chain.Blocks)-1]
}

// adds a new block to the blockchain
func (chain *BlockChain) pushBlock(block *Block) {
	chain.Blocks = append(chain.Blocks, block)
}

// generates and adds a new block to the blockchain
func (chain *BlockChain) GenerateBlock(data string) {
	head := chain.Head()
	next := CreateBlock(data, head.Hash)
	chain.pushBlock(next)
}

// static

// initial block for any blockchain
func Genesis() *Block {
	return CreateBlock("GENESIS", []byte{})
}

// creates a new blockchain
func CreateBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{Genesis()},
	}
}
