package prompts

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/t02smith/basic-blockchain/blockchain"
	"github.com/t02smith/basic-blockchain/wallet"
)

// prompts a user to choose between the addresses in their wallet
func PromptAddress(text string) (string, error) {
	wallets, _ := wallet.CreateWallet()
	addrs := wallets.GetAllAddresses()
	if len(addrs) == 0 {
		return "", errors.New("you must create a wallet first")
	}

	bcExists := blockchain.BlockchainExists()

	items := []string{}
	for _, addr := range addrs {
		w := wallets.Wallets[addr]

		if !bcExists {
			if len(w.Alias) > 0 {
				items = append(items, fmt.Sprintf("%s %s", w.Alias, addr))
			} else {
				items = append(items, addr)
			}

			continue
		}

		bal := wallet.GetBalance(addr)
		if len(w.Alias) > 0 {
			items = append(items, fmt.Sprintf("%s %s => %d", w.Alias, addr, bal))
		} else {
			items = append(items, fmt.Sprintf("%s => %d", addr, bal))
		}
	}

	prompt := promptui.Select{
		Label: text,
		Items: items,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Println("Failed to choose address.")
		return "", err
	}

	return addrs[i], nil
}

// Prompt the user to choose an address but only show addresses that satisfy a predicate
func PromptAddressIf(text string, predicate func(address string) bool) (string, error) {
	if !blockchain.BlockchainExists() {
		return "", errors.New("no blockchain found")
	}

	wallets, err := wallet.CreateWallet()
	if err != nil {
		return "", err
	}

	addrs, items, keptAddrs := wallets.GetAllAddresses(), []string{}, []string{}

	for _, addr := range addrs {
		if predicate(addr) {
			w := wallets.Wallets[addr]
			keptAddrs = append(keptAddrs, addr)
			bal := wallet.GetBalance(addr)
			if len(w.Alias) > 0 {
				items = append(items, fmt.Sprintf("%s %s => %d", w.Alias, addr, bal))
			} else {
				items = append(items, fmt.Sprintf("%s => %d", addr, bal))
			}
		}
	}

	prompt := promptui.Select{
		Label: text,
		Items: items,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Println("Failed to choose address.")
		return "", err
	}

	return keptAddrs[i], nil

}

// Prompt the user to choose an amount of coin to send from an address
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
