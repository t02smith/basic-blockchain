package wallet

import "crypto/ecdsa"

// a wallet consists of a public/private key pair
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
	Alias      string
}

// each user may have many wallets for privacy
type Wallets struct {
	Wallets map[string]*Wallet
}
