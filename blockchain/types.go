package blockchain

import (
	"math/big"

	"github.com/dgraph-io/badger"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}
