package cli

import (
	"fmt"
	"strconv"

	"github.com/t02smith/basic-blockchain/blockchain"
)

func (cli *CommandLine) printUsage() {
	fmt.Println("Commands:")
	fmt.Println("getbalance -address ADDRESS = get balance for ADDRESS")
	fmt.Println("createblockchain -address ADDRESS creates a blockchain and rewards the mining fee")
	fmt.Println("printchain - Prints the blocks in the chain")
	fmt.Println("send -from FROM -to TO -amount AMOUNT - Send amount of coins from one address to another")
}

func (cli *CommandLine) createBlockChain(address string) {
	newChain := blockchain.CreateBlockChain(address)
	newChain.Database.Close()
	fmt.Println("Created new BlockChain")
}

func (cli *CommandLine) getBalance(address string) {
	chain := blockchain.ContinueBlockChain(address)
	defer chain.Database.Close()

	balance := 0
	UTXOs := chain.FindUTXOs(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of %s: %d\n", address, balance)
}

func (cli *CommandLine) send(from, to string, amount int) {
	chain := blockchain.ContinueBlockChain(from)
	defer chain.Database.Close()

	txn := blockchain.NewTransaction(from, to, amount, chain)
	chain.GenerateBlock([]*blockchain.Transaction{txn})
	fmt.Println("Success!")
}

func (cli *CommandLine) printChain() {
	chain := blockchain.ContinueBlockChain("")
	defer chain.Database.Close()

	it := chain.Iterator()

	for {
		block := it.Next()
		fmt.Printf("%x -> %x:\n", block.PrevHash, block.Hash)
		pow := blockchain.CreateProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(blockchain.ValidateProofOfWork(pow)))

		if len(block.PrevHash) == 0 {
			break
		}
	}
}
