package prompts

import (
	"errors"
	"fmt"
	"strconv"

	mapset "github.com/deckarep/golang-set"
	"github.com/manifoldco/promptui"
	"github.com/t02smith/basic-blockchain/wallet"
)

// prompts a user to choose between the addresses in their wallet
func PromptAddress(text string) (string, error) {
	wallets, _ := wallet.CreateWallet()
	if len(wallets.GetAllAddresses()) == 0 {
		return "", errors.New("you must create a wallet first")
	}

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

func PromptAddressButExclude(text string, exclude []string) (string, error) {
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
func PromptForTxnAmount(max int) (int, error) {
	prompt, _ := IntegerPrompt("Enter an amount to send", 0, max)

	result, err := prompt.Run()
	if err != nil {
		return -1, err
	}

	number, err := strconv.ParseInt(result, 10, 32)
	if err != nil {
		return -1, err
	}

	return int(number), nil
}
