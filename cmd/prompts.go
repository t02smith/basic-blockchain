package cmd

import (
	"errors"
	"fmt"
	"strconv"

	mapset "github.com/deckarep/golang-set"
	"github.com/manifoldco/promptui"
	"github.com/t02smith/basic-blockchain/wallet"
)

// prompts a user to choose between the addresses in their wallet
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

func promptAddressButExclude(text string, exclude []string) (string, error) {
	wallets, _ := wallet.CreateWallet()

	addresses := mapset.NewSet()
	for _, addr := range wallets.GetAllAddresses() {
		addresses.Add(addr)
	}

	excludeSet := mapset.NewSet()
	for _, excl := range exclude {
		excludeSet.Add(excl)
	}

	addressArr := addresses.Difference(excludeSet).ToSlice()

	prompt := promptui.Select{
		Label: text,
		Items: addressArr,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Failed to choose address.")
		return "", err
	}

	return result, nil
}

//

func numberPrompt(label string, min, max int) (*promptui.Prompt, error) {
	if min >= max {
		return &promptui.Prompt{}, errors.New("min must be smaller than max")
	}

	return &promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			res, err := strconv.ParseInt(s, 10, 32)
			x := int(res)
			if err != nil || x < min || x > max {
				return errors.New("invalid number")
			}
			return nil
		},
	}, nil
}

func promptForTxnAmount(max int) int {
	prompt, _ := numberPrompt("Enter an amount to send: ", 0, max)

	result, _ := prompt.Run()
	number, _ := strconv.ParseInt(result, 10, 32)
	return int(number)
}
