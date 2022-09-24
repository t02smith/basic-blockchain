package store

import (
	"fmt"
	"os"
)

func InitStore(location string) *Store {
	if _, err := os.Stat(location); !os.IsNotExist(err) {
		// read store
	}

	os.Mkdir(location, os.FileMode(777))
	return &Store{
		Location: location,
		Shards:   []*Shard{},
	}
}

func readStore(location string) (*Store, error) {
	var shards []*Shard

	dir, err := os.ReadDir(location)
	if err != nil {
		fmt.Printf("Error reading %s\n", location)
		return &Store{}, fmt.Errorf("error reading %s", location)
	}

	for _, item := range dir {
		if item.IsDir() {
			continue
		}

		shards = append(shards, &Shard{})
	}

	return &Store{
		Location: location,
		Shards:   shards,
	}, nil
}
