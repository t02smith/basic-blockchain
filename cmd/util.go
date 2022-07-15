package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/t02smith/basic-blockchain/wallet"
)

func promptAddress(text string) (string, error) {
	wallets, _ := wallet.CreateWallet()

	prompt := promptui.Select{
		Label: text,
		Items: wallets.GetAllAddresses(),
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Failed to choose address.")
		return "", err
	}

	return result, nil
}

func promptForTxnAmount() int {
	prompt := promptui.Prompt{
		Label: "Enter an amount to send: ",
		Validate: func(s string) error {
			_, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return errors.New("invalid number")
			}
			return nil
		},
	}

	result, _ := prompt.Run()
	number, _ := strconv.ParseInt(result, 10, 64)
	return int(number)
}
