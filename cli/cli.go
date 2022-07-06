package cli

import (
	"flag"
	"os"
	"runtime"

	"github.com/t02smith/basic-blockchain/blockchain"
)

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}

		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

}
