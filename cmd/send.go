package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/blockchain"
)

var (
	from, to string
	amount   int
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send currency from one account to another",
	Long:  `Send unspent currency from one account to another providing the sender has sufficient funds.`,
	Run: func(cmd *cobra.Command, args []string) {

		if from == "" {
			from, _ = promptAddress("Choose an address to send from:")
		}

		if to == "" {
			to, _ = promptAddress("Choose an address to send to:")
		}

		if amount == -1 {
			amount = promptForTxnAmount()
		}

		send(from, to, amount)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.PersistentFlags().StringVarP(&from, "from", "f", "", "Address to send the currency from")
	sendCmd.PersistentFlags().StringVarP(&to, "to", "t", "", "Address to send the currency to")

	sendCmd.PersistentFlags().IntVarP(&amount, "amount", "a", -1, "Amount of currency to send")
}

func send(from, to string, amount int) {
	chain := blockchain.ContinueBlockChain(from)
	defer chain.Database.Close()

	txn := blockchain.NewTransaction(from, to, amount, chain)
	chain.GenerateBlock([]*blockchain.Transaction{txn})
	fmt.Println("Success!")
}
