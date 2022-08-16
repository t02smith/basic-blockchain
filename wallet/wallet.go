package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"github.com/t02smith/basic-blockchain/blockchain"
)

const (
	checksumLength int = 4

	version byte = byte(0x00)
)

// GETTERS

// generates the address of a wallet
func (w *Wallet) address() []byte {
	pubHash := publicKeyHash(w.PublicKey)
	versionedHash := append([]byte{version}, pubHash...)

	checksum := checksum(versionedHash)
	finalHash := append(versionedHash, checksum...)

	return base58Encode(finalHash)
}

// CREATORS

// creates a new wallet
func makeWallet(alias string) *Wallet {
	private, public := newKeyPair()
	return &Wallet{
		PrivateKey: private,
		PublicKey:  public,
		Alias:      alias,
	}
}

// generates a new public/private key pair
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panicln("Error generating private key.")
	}

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pub
}

// returns the hash of a given public key
func publicKeyHash(publicKey []byte) []byte {
	hashedPublicKey := sha256.Sum256(publicKey)

	hasher := sha256.New()
	_, err := hasher.Write(hashedPublicKey[:])
	if err != nil {
		log.Panicln(err)
	}

	return hasher.Sum(nil)
}

// returns the checksum of a versioned hash
func checksum(ripeMdHash []byte) []byte {
	firstHash := sha256.Sum256(ripeMdHash)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]
}

// gets the balance at a given address
func GetBalance(address string) int {
	chain := blockchain.ContinueBlockChain(address)
	defer chain.Database.Close()

	balance := 0
	UTXOs := chain.FindUTXOs(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	return balance
}
