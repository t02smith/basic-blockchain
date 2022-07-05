package main

import (
	"fmt"

	"github.com/t02smith/basic-blockchain/blockchain"
)

func main() {
	chain := blockchain.CreateBlockChain()

	chain.GenerateBlock("hello world")
	chain.GenerateBlock("boobies")

	for _, block := range chain.Blocks {
		fmt.Printf("%x -> %x:\n %s\n\n", block.PrevHash, block.Hash, block.Data)
	}
}
