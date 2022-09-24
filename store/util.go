package store

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func encrypt(data, secret []byte) ([]byte, error) {

	block, err := aes.NewCipher(secret[:32])
	if err != nil {
		return []byte{}, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, err
	}

	cipherText := aesGCM.Seal(nonce, nonce, data, nil)
	return cipherText, nil

}
