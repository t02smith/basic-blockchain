package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"github.com/t02smith/basic-blockchain/blockchain"
	"golang.org/x/crypto/ripemd160"
)

const (
	checksumLength int = 4

	version byte = byte(0x00)
)

func (w *Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PublicKey)
	versionedHash := append([]byte{version}, pubHash...)

	checksum := Checksum(versionedHash)
	finalHash := append(versionedHash, checksum...)

	return base58Encode(finalHash)
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	return &Wallet{
		PrivateKey: private,
		PublicKey:  public,
	}
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panicln("Error generating private key.")
	}

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pub
}

func PublicKeyHash(publicKey []byte) []byte {
	hashedPublicKey := sha256.Sum256(publicKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(hashedPublicKey[:])
	if err != nil {
		log.Panicln(err)
	}

	return hasher.Sum(nil)
}

func Checksum(ripeMdHash []byte) []byte {
	firstHash := sha256.Sum256(ripeMdHash)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]
}

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
