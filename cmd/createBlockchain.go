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
	Short: "create a new blockchain",
	Long:  `Create a new blockchain and award a given address for the creation of the genesis block`,
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
