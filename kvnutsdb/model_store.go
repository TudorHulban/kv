package kvnuts

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nutsdb/nutsdb"
)

type KV struct {
	Key   []byte
	Value []byte
}

type KVStore struct {
	Store *nutsdb.DB

	bucket string
}

// NewStoreInMemory returns a type containing a store that satisfies store interface.
// With test segment size.
func NewStoreInMemory(mbSegmentSize uint) (*KVStore, error) {
	db, errOpen := nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir(_folderDB+strconv.Itoa(int(time.Now().UnixNano()))),
		nutsdb.WithSegmentSize(int64(mbSegmentSize)),
	)
	if errOpen != nil {
		return nil,
			fmt.Errorf("could not create database in folder: %s, %w",
				_folderDB, errOpen)
	}

	return &KVStore{
			Store: db,
		},
		nil
}

// NewStore returns a type containing a store that satisfies store interface.
// With test segment size.
func NewStore(mbSegmentSize uint) (*KVStore, error) {
	db, errOpen := nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir(_folderDB),
		nutsdb.WithEntryIdxMode(nutsdb.HintKeyAndRAMIdxMode),
		nutsdb.WithSegmentSize(int64(mbSegmentSize)),
	)
	if errOpen != nil {
		return nil,
			fmt.Errorf("could not create database in folder: %s, %w",
				_folderDB, errOpen)
	}

	return &KVStore{
			Store: db,
		},
		nil
}

// Close closes the store.
func (s *KVStore) Close() error {
	return s.Store.Close()
}
