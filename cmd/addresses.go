/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/wallet"
)

// addressesCmd represents the addresses command
var addressesCmd = &cobra.Command{
	Use:   "addresses",
	Short: "List all stored addresses",
	Long:  `List all stored addresses currently in your wallet`,
	Run: func(cmd *cobra.Command, args []string) {
		listAddresses()
	},
}

func init() {
	rootCmd.AddCommand(addressesCmd)

}

func listAddresses() {
	wallets, _ := wallet.CreateWallet()
	addresses := wallets.GetAllAddresses()

	for i, address := range addresses {
		balance := wallet.GetBalance(address)
		fmt.Printf("%d. %s = %d\n", i+1, address, balance)
	}
}
