package blockchain

import (
	"log"

	"github.com/dgraph-io/badger"
)

const DB_PATH string = "./tmp/blocks"

// methods

// generates and adds a new block to the blockchain
// func (chain *BlockChain) GenerateBlock(data string) {
// 	log.Printf("Generating block: %s\n", data)
// 	head := chain.Head()
// 	next := CreateBlock(data, head.Hash)
// 	chain.pushBlock(next)
// }

func (chain *BlockChain) GenerateBlock(data string) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		Handle(err)
		return err
	})
	Handle(err)

	newBlock := CreateBlock(data, lastHash)
	err = chain.Database.Update(func(txn *badger.Txn) error {
		err = txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash
		return err
	})
	Handle(err)
}

// static

// initial block for any blockchain
func Genesis() *Block {
	return CreateBlock("GENESIS", []byte{})
}

// creates a new blockchain
func CreateBlockChain() *BlockChain {
	// open connection to db
	var prevHash []byte
	ops := badger.DefaultOptions(DB_PATH)
	ops.Logger = nil

	db, err := badger.Open(ops)
	Handle(err)

	// look for existing blockchain
	err = db.Update(func(txn *badger.Txn) error {

		// create new database if one doesn't exist
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			log.Println("No existing blockchain found")

			genesis := Genesis()
			log.Println("Generated GENESIS")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)

			err = txn.Set([]byte("lh"), genesis.Hash)
			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			err = item.Value(func(val []byte) error {
				prevHash = val
				return nil
			})
			Handle(err)
			return err
		}
	})
	Handle(err)

	blockchain := BlockChain{
		LastHash: prevHash,
		Database: db,
	}
	return &blockchain
}

// ITERATOR

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iterator := BlockChainIterator{
		CurrentHash: chain.LastHash,
		Database:    chain.Database,
	}

	return &iterator
}

func (it *BlockChainIterator) Next() *Block {
	var b *Block

	err := it.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(it.CurrentHash)
		Handle(err)

		err = item.Value(func(val []byte) error {
			b = Deserialize(val)
			return nil
		})
		Handle(err)
		return err
	})
	Handle(err)

	it.CurrentHash = b.PrevHash
	return b
}
