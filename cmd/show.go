package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/blockchain"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Print the blockchain",
	Long:  `Print all the blocks in the blockchain`,
	Run: func(cmd *cobra.Command, args []string) {
		printChain()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

}

func printChain() {
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
