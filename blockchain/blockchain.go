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
		handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		handle(err)
		return err
	})
	handle(err)

	newBlock := createBlock(transactions, lastHash)
	err = chain.Database.Update(func(txn *badger.Txn) error {
		err = txn.Set(newBlock.Hash, newBlock.serialize())
		handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash
		return err
	})
	handle(err)
}

// static

// initial block for any blockchain
func genesis(coinbase *Transaction) *Block {
	return createBlock([]*Transaction{coinbase}, []byte{})
}

func CreateBlockChain(address string) *BlockChain {
	var lastHash []byte

	if BlockchainExists() {
		log.Println("Blockchain already exists")
		runtime.Goexit()
	}

	ops := badger.DefaultOptions(DB_PATH)
	ops.Logger = nil

	db, err := badger.Open(ops)
	handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		cbtxn := CoinbaseTxn(address, GENESIS_DATA)
		genesis := genesis(cbtxn)
		log.Println("Genesis created")

		err = txn.Set(genesis.Hash, genesis.serialize())
		handle(err)

		err = txn.Set([]byte("lh"), genesis.Hash)
		lastHash = genesis.Hash

		return err
	})
	handle(err)

	chain := BlockChain{
		LastHash: lastHash,
		Database: db,
	}

	return &chain
}

func ContinueBlockChain(address string) *BlockChain {
	if !BlockchainExists() {
		fmt.Println("No blockchain found, please create one first")
		runtime.Goexit()
	}

	var lastHash []byte

	opts := badger.DefaultOptions(DB_PATH)
	opts.Logger = nil
	db, err := badger.Open(opts)
	handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		handle(err)
		return err
	})
	handle(err)

	chain := BlockChain{lastHash, db}
	return &chain
}

func BlockchainExists() bool {
	_, err := os.Stat(DB_FILE)
	return !os.IsNotExist(err)
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
		handle(err)

		err = item.Value(func(val []byte) error {
			b = deserialize(val)
			return nil
		})
		handle(err)
		return err
	})
	handle(err)

	it.CurrentHash = b.PrevHash
	return b
}

// HASHING

func (b *Block) hashTransactions() []byte {
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
func (chain *BlockChain) findUnspentTxns(address string) []Transaction {
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
	unspent := chain.findUnspentTxns(address)

	for _, tx := range unspent {
		for _, out := range tx.Outputs {
			if out.CanBeUnlocked(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

func (chain *BlockChain) findSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOuts := make(map[string][]int)
	unspentTxns := chain.findUnspentTxns(address)
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
