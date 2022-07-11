package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
)

// CONSTANTS

const (
	DB_PATH      string = "./tmp/blocks"
	DB_FILE      string = "./tmp/blocks/MANIFEST"
	GENESIS_DATA string = "TXN #1"
)

// FUNCTIONS

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panicln(err)
	}

	return buff.Bytes()
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
