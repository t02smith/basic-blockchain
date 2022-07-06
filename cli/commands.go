package cli

import (
	"fmt"
	"strconv"

	"github.com/t02smith/basic-blockchain/blockchain"
)

func (cli *CommandLine) printUsage() {
	fmt.Println("Commands:")
	fmt.Println(" add -block <BLOCK DATA>  - add a block to the chain")
	fmt.Println(" print 				   - prints the blocks in the chain")
}

func (cli *CommandLine) addBlock(data string) {
	cli.Blockchain.GenerateBlock(data)
	fmt.Printf("Block added: %s\n", data)
}

func (cli *CommandLine) printChain() {
	it := cli.Blockchain.Iterator()

	for {
		block := it.Next()
		fmt.Printf("%x -> %x:\n %s\n", block.PrevHash, block.Hash, block.Data)
		pow := blockchain.CreateProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(blockchain.ValidateProofOfWork(pow)))

		if len(block.PrevHash) == 0 {
			break
		}
	}
}
