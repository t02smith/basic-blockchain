package main

import (
	"os"

	"github.com/t02smith/basic-blockchain/blockchain"
	"github.com/t02smith/basic-blockchain/cli"
)

func main() {
	defer os.Exit(0)

	chain := blockchain.CreateBlockChain()
	defer chain.Database.Close()

	cli := cli.CommandLine{
		Blockchain: chain,
	}

	cli.Run()
}
