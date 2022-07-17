/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/wallet"
)

var walletAmount int

// createWalletCmd represents the createWallet command
var createWalletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Create a new wallet",
	Long:  `Create a new wallet and add it to your collection of wallets`,
	Run: func(cmd *cobra.Command, args []string) {
		if walletAmount < 1 {
			prompt := promptui.Prompt{
				Label: "Enter a valid number of wallets to create (>= 1)",
				Validate: func(s string) error {
					v, err := strconv.ParseInt(s, 10, 64)
					if err != nil || v < 1 {
						return errors.New("invalid number")
					}

					return nil
				},
			}

			result, _ := prompt.Run()
			number, _ := strconv.ParseInt(result, 10, 64)
			walletAmount = int(number)
		}

		for i := 0; i < walletAmount; i++ {
			createWallet()
		}

	},
}

func init() {
	createCmd.AddCommand(createWalletCmd)

	createWalletCmd.PersistentFlags().IntVarP(&walletAmount, "amount", "n", 1, "The amount of new wallets to make")
}

func createWallet() {
	wallets, _ := wallet.CreateWallet()
	address := wallets.AddWallet()
	wallets.SaveFile()

	fmt.Printf("Wallet created with address %s\n", address)
}
