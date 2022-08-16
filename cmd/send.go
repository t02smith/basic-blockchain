package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/basic-blockchain/blockchain"
	"github.com/t02smith/basic-blockchain/cmd/prompts"
	"github.com/t02smith/basic-blockchain/wallet"
)

var (
	from, to string
	amount   int
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send currency from one account to another",
	Long: `Send unspent currency from one account to another providing the sender has sufficient funds.
	A prompt will show up to help you choose the to and from addresses as well as the amount`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		if from == "" {
			from, err = prompts.PromptAddressIf("Choose an address to send from:", func(address string) bool {
				bal := wallet.GetBalance(address)
				return bal > 0
			})
			if err != nil {
				fmt.Println("Error while choosing origin address.")
				return
			}
		}

		if to == "" {
			to, err = prompts.PromptAddressIf("Choose an address to send to:", func(address string) bool {
				return address != from
			})
			if err != nil {
				fmt.Println("Error choosing destination address.")
				return
			}
		}

		bal := wallet.GetBalance(from)
		if amount == -1 {
			amount, err = prompts.PromptForTxnAmount(bal)
			if err != nil {
				fmt.Println("Error selecting amount to send.")
				return
			}
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
