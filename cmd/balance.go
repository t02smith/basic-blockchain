package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/wallet"
)

var balanceAddress string

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Get the balance at an address",
	Long:  `Finds the total amount of unspent output from a given address' transactions.`,
	Run: func(cmd *cobra.Command, args []string) {
		if balanceAddress == "" {
			balanceAddress, _ = promptAddress("Choose an address to check the balance of:")
		}

		getBalance(balanceAddress)
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	balanceCmd.PersistentFlags().StringVarP(&balanceAddress, "address", "a", "", "Address to check balance of")
}

func getBalance(address string) {
	balance := wallet.GetBalance(address)

	fmt.Printf("Balance of %s: %d\n", address, balance)
}
