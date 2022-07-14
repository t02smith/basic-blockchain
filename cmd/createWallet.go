/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/wallet"
)

// createWalletCmd represents the createWallet command
var createWalletCmd = &cobra.Command{
	Use:   "createWallet",
	Short: "Create a new wallet",
	Long:  `Create a new wallet and add it to your collection of wallets`,
	Run: func(cmd *cobra.Command, args []string) {
		createWallet()
	},
}

func init() {
	rootCmd.AddCommand(createWalletCmd)

}

func createWallet() {
	wallets, _ := wallet.CreateWallet()
	address := wallets.AddWallet()
	wallets.SaveFile()

	fmt.Printf("Wallet created with address %s\n", address)
}
