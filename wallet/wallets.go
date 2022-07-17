package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
)

const WALLET_FILE string = "./tmp/wallets.data"

// database

func (ws *Wallets) SaveFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panicln(err)
	}

	err = ioutil.WriteFile(WALLET_FILE, content.Bytes(), 0644)
	if err != nil {
		log.Panicln(err)
	}
}

func (ws *Wallets) LoadFile() error {
	if _, err := os.Stat(WALLET_FILE); os.IsNotExist(err) {
		return err
	}

	var wallets Wallets
	fileContent, err := ioutil.ReadFile(WALLET_FILE)
	if err != nil {
		return err
	}

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		return err
	}

	ws.Wallets = wallets.Wallets
	return nil
}

func CreateWallet() (*Wallets, error) {
	wallets := Wallets{
		Wallets: make(map[string]*Wallet),
	}

	err := wallets.LoadFile()
	return &wallets, err
}

func (ws *Wallets) AddWallet() string {
	wallet := MakeWallet()
	address := string(wallet.Address())

	ws.Wallets[address] = wallet
	return address
}

// getters

func (ws *Wallets) GetWallet(address string) *Wallet {
	return ws.Wallets[address]
}

func (ws *Wallets) GetAllAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

func (ws *Wallets) GetAllAddressesAndBalance() map[string]int {
	pairs := make(map[string]int)
	for _, addr := range ws.GetAllAddresses() {
		pairs[addr] = GetBalance(addr)
	}

	return pairs
}
