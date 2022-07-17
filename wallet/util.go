package wallet

import (
	"github.com/mr-tron/base58/base58"
)

func base58Encode(input []byte) []byte {
	return []byte(base58.Encode(input))
}
