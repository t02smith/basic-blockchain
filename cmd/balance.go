package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/blockchain"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Get the balance at an address",
	Long:  `Finds the total amount of unspent output from a given address' transactions.`,
	Run: func(cmd *cobra.Command, args []string) {
		getBalance(address)
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	balanceCmd.PersistentFlags().StringVarP(&address, "address", "a", "", "Address to check balance of")
	balanceCmd.MarkPersistentFlagRequired("address")
}

func getBalance(address string) {
	chain := blockchain.ContinueBlockChain(address)
	defer chain.Database.Close()

	balance := 0
	UTXOs := chain.FindUTXOs(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of %s: %d\n", address, balance)
}
