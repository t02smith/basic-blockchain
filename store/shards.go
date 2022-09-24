package store

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// TODO store shard data in db
// TODO read shards

func ShardFile(location, output string, publicKey []byte) error {
	return shardFile(location, output, publicKey, SHARD_SIZE)
}

func shardFile(location, output string, publicKey []byte, shardSize int) error {
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return err
	}

	if item, err := os.Stat(output); os.IsNotExist(err) || !item.IsDir() {
		return fmt.Errorf("error finding ouput folder %s", output)
	}

	file, err := os.Open(location)
	if err != nil {
		return err
	}

	buffer := make([]byte, shardSize)
	reader := bufio.NewReader(file)

	fmt.Printf("Sharding file '%s'\n", location)
	for {
		_, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		hash := sha256.Sum256(buffer)
		encyrpted, err := encrypt(buffer, publicKey)
		if err != nil {
			fmt.Println(err)
			continue
		}

		ioutil.WriteFile(fmt.Sprintf("%s/%x.shard", output, hash), encyrpted, 0644)

	}

	return nil
}
