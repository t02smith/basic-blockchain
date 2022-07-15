/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/blockchain"
)

var createAddress string = ""

// createBlockchainCmd represents the createBlockchain command
var createBlockchainCmd = &cobra.Command{
	Use:   "blockchain",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if createAddress == "" {
			createAddress, _ = promptAddress("Choose an address to award for the creation of the genesis block")
		}

		createBlockChain(createAddress)
	},
}

func init() {
	createCmd.AddCommand(createBlockchainCmd)

	createBlockchainCmd.PersistentFlags().StringVarP(&createAddress, "address", "a", "", "Address to send initial reward to")
}

func createBlockChain(address string) {
	newChain := blockchain.CreateBlockChain(address)
	newChain.Database.Close()
	fmt.Println("Created new BlockChain")
}
