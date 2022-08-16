/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/blockchain"
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

	i := 1
	if blockchain.BlockchainExists() {
		for addr, w := range wallets.Wallets {
			balance := wallet.GetBalance(addr)
			fmt.Printf("%d. %s - %s => %d\n", i, w.Alias, addr, balance)
			i++
		}
	} else {
		fmt.Println("WARNING - No blockchain found")
		for addr, w := range wallets.Wallets {
			fmt.Printf("%d. %s - %s\n", i, w.Alias, addr)
			i++
		}
	}

}
