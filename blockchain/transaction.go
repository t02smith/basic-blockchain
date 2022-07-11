package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

const REWARD int = 100

// UTIL

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func (in *TxnInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxnOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

func (txn *Transaction) IsCoinbase() bool {
	return len(txn.Inputs) == 1 && len(txn.Inputs[0].ID) == 0 && txn.Inputs[0].Out == -1
}

//

func CoinbaseTxn(toAddress, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", toAddress)
	}

	txIn := TxnInput{
		ID:  []byte{},
		Out: -1,
		Sig: data,
	}

	txOut := TxnOutput{
		Value:  REWARD,
		PubKey: toAddress,
	}

	return &Transaction{
		ID:      nil,
		Inputs:  []TxnInput{txIn},
		Outputs: []TxnOutput{txOut},
	}
}

func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxnInput
	var outputs []TxnOutput

	// find spendable outputs
	acc, validOutputs := chain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panicln("Error: Not enough funds")
	}

	// generate inputs that point to the outputs being spent
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxnInput{
				ID:  txID,
				Out: out,
				Sig: from,
			}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxnOutput{
		Value:  amount,
		PubKey: to,
	})

	// store any leftover currency
	if acc > amount {
		outputs = append(outputs, TxnOutput{
			Value:  acc - amount,
			PubKey: from,
		})
	}

	txn := Transaction{
		ID:      nil,
		Inputs:  inputs,
		Outputs: outputs,
	}
	txn.SetID()

	return &txn

}
