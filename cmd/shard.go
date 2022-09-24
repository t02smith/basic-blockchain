/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/store"
	"github.com/t02smith/basic-blockchain/wallet"
)

var (
	shardAddress, output string
)

// shardCmd represents the shard command
var shardCmd = &cobra.Command{
	Use:   "shard",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ws, err := wallet.CreateWallet()
		if err != nil {
			fmt.Println(err)
			return
		}

		pubKey, err := ws.GetPublicKey(shardAddress)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = store.ShardFile(args[0], output, pubKey)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(shardCmd)

	shardCmd.Flags().StringVarP(&output, "output", "o", "", "Where to store shards")
	shardCmd.Flags().StringVarP(&shardAddress, "address", "a", "", "Public key to use to encrypt shards")
}
