/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/blockchain"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new blockchain",
	Long:  `Create a new blockchain and send the reward for mining the genesis block the given address.`,
	Run: func(cmd *cobra.Command, args []string) {
		createBlockChain(address)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().StringVarP(&address, "address", "a", "", "Address to send initial reward to")
	createCmd.MarkPersistentFlagRequired("address")
}

func createBlockChain(address string) {
	newChain := blockchain.CreateBlockChain(address)
	newChain.Database.Close()
	fmt.Println("Created new BlockChain")
}
