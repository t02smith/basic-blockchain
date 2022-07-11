package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/dgraph-io/badger"
)

// methods

func (chain *BlockChain) GenerateBlock(transactions []*Transaction) {
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

	newBlock := CreateBlock(transactions, lastHash)
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
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

func DbExists(db string) bool {
	if _, err := os.Stat(db); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateBlockChain(address string) *BlockChain {
	var lastHash []byte

	if DbExists(DB_FILE) {
		log.Println("Blockchain already exists")
		runtime.Goexit()
	}

	ops := badger.DefaultOptions(DB_PATH)
	ops.Logger = nil

	db, err := badger.Open(ops)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		cbtxn := CoinbaseTxn(address, GENESIS_DATA)
		genesis := Genesis(cbtxn)
		log.Println("Genesis created")

		err = txn.Set(genesis.Hash, genesis.Serialize())
		Handle(err)

		err = txn.Set([]byte("lh"), genesis.Hash)
		lastHash = genesis.Hash

		return err
	})
	Handle(err)

	chain := BlockChain{
		LastHash: lastHash,
		Database: db,
	}

	return &chain
}

func ContinueBlockChain(address string) *BlockChain {
	if !DbExists(DB_FILE) {
		fmt.Println("No blockchain found, please create one first")
		runtime.Goexit()
	}

	var lastHash []byte

	opts := badger.DefaultOptions(DB_PATH)
	opts.Logger = nil
	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
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

	chain := BlockChain{lastHash, db}
	return &chain
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

// HASHING

func (b *Block) HashTransactions() []byte {
	var txnHashes [][]byte
	var txnHash [32]byte

	for _, txn := range b.Transactions {
		txnHashes = append(txnHashes, txn.ID)
	}

	txnHash = sha256.Sum256(bytes.Join(txnHashes, []byte{}))
	return txnHash[:]
}

// UXTOs

// finds all unspent txns for a given address
func (chain *BlockChain) FindUnspentTxns(address string) []Transaction {
	var uxtos []Transaction

	spentTXNs := make(map[string][]int)
	it := chain.Iterator()

	for { // loop through entire blockchain from most recent -> oldest block
		block := it.Next()

		// loop through each block's transactions
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			// loop through a transaction's outputs
			for outIDx, out := range tx.Outputs {

				// the output has been spent
				if spentTXNs[txID] != nil {
					for _, spentOut := range spentTXNs[txID] {
						if spentOut == outIDx {
							continue Outputs
						}
					}
				}

				// this uxto belongs to the given address
				if out.CanBeUnlocked(address) {
					uxtos = append(uxtos, *tx)
				}
			}

			// if not coinbase, loop through inputs
			// if the address matches, we know it's spent
			if !tx.IsCoinbase() {
				for _, in := range tx.Inputs {
					if in.CanUnlock(address) {
						inTxnID := hex.EncodeToString(in.ID)
						spentTXNs[inTxnID] = append(spentTXNs[inTxnID], in.Out)
					}
				}
			}
		}

		if len(block.PrevHash) == 0 {
			break
		}
	}

	return uxtos
}

// finds UTXOs related to an address
func (chain *BlockChain) FindUTXOs(address string) []TxnOutput {
	var UTXOs []TxnOutput
	unspent := chain.FindUnspentTxns(address)

	for _, tx := range unspent {
		for _, out := range tx.Outputs {
			if out.CanBeUnlocked(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

func (chain *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOuts := make(map[string][]int)
	unspentTxns := chain.FindUnspentTxns(address)
	accumulated := 0

Work:
	for _, txn := range unspentTxns {
		txID := hex.EncodeToString(txn.ID)
		for outIdx, out := range txn.Outputs {
			if out.CanBeUnlocked(address) && accumulated < amount {
				accumulated += out.Value
				unspentOuts[txID] = append(unspentOuts[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOuts
}
