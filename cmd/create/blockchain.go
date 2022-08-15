/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package create

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/blockchain"
	"github.com/t02smith/basic-blockchain/cmd/prompts"
)

var createAddress string = ""

// createBlockchainCmd represents the createBlockchain command
var createBlockchainCmd = &cobra.Command{
	Use:   "blockchain",
	Short: "create a new blockchain",
	Long:  `Create a new blockchain and award a given address for the creation of the genesis block`,
	Run: func(cmd *cobra.Command, args []string) {
		if createAddress == "" {
			var err error
			createAddress, err = prompts.PromptAddress("Choose an address to award for the creation of the genesis block")
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		createBlockChain(createAddress)
	},
}

func init() {
	CreateCmd.AddCommand(createBlockchainCmd)

	createBlockchainCmd.PersistentFlags().StringVarP(&createAddress, "address", "a", "", "Address to send initial reward to")
}

func createBlockChain(address string) {
	newChain := blockchain.CreateBlockChain(address)
	newChain.Database.Close()
	fmt.Println("Created new BlockChain")
}
