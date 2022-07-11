package blockchain

import (
	"math/big"

	"github.com/dgraph-io/badger"
)

// BLOCK

type Block struct {
	Hash         []byte
	Transactions []*Transaction
	PrevHash     []byte
	Nonce        int
}

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// BLOCKCHAIN

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// TRANSACTIONS

type Transaction struct {
	ID      []byte
	Inputs  []TxnInput
	Outputs []TxnOutput
}

type TxnInput struct {
	ID  []byte // finds the txn that a given output is in
	Out int    // index of the specific output if a txn has many outputs
	Sig string // script that adds data to an outputs PubKey
}

type TxnOutput struct {
	Value  int    // coins transferred
	PubKey string // uniquely identifies a user
}
