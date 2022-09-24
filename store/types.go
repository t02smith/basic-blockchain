package store

const (
	SHARD_SIZE int = 256000 // bytes
)

type Store struct {

	// filepath to local store
	Location string

	// shards stored in store
	Shards []*Shard
}

// Stores an encrypted section of an object
type Shard struct {

	// Unique identifier of shard
	Hash []byte

	// Data stored in shard
	Data []byte
}
