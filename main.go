package main

import (
	"os"

	"github.com/t02smith/basic-blockchain/cli"
)

func main() {
	defer os.Exit(0)

	cli := cli.CommandLine{}
	cli.Run()
}
