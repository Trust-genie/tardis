package storage

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

func init() {
	Store = *NewStore(10, 20)

}

var (
	Store Storage
)

// A storage is a Sharded map table that is particularly useful as a
// high read, low lock contention in memory database
// it can recieve and store oll forms of datatypes not just primitives
// and shines in storing go struct types.
type Storage []*shard

// we use vertical sharding to reduce the possiblity of lock contentions
type shard struct {
	m map[string]interface{}
	sync.RWMutex
}

func NewStore(Shardno, Shardcap uint) *Storage {
	var store Storage = make([]*shard, Shardno)

	for i := 0; i < int(Shardno); i++ {
		store[i] = &shard{
			m: make(map[string]interface{}, Shardcap),
		}
	}
	return &store
}

func (m Storage) getShard(key string) *shard {
	checksum := sha1.Sum([]byte(key))
	hash := int(checksum[10])

	shardno := hash % len(m)

	return m[shardno]

}

func (m Storage) getShardId(key string) int {
	checksum := sha1.Sum([]byte(key))
	hash := int(checksum[10])
	shardno := hash % len(m)

	return shardno
}

// Put is a safe way to put values into our database
//
// it returns an error if there is a an existing key.
// For updates use Insert
func (s *Storage) Put(key string, value interface{}) error {
	_, err := s.Get(key)
	if err == nil {
		return fmt.Errorf("Key Exist, %s", err)
	}

	store := s.getShard(key)

	store.Lock()
	defer store.Unlock()

	store.m[key] = value

	return nil

}

// Retireves data from Database
func (s *Storage) Get(key string) (interface{}, error) {

	store := s.getShard(key)
	store.RLock()
	value, ok := store.m[key]
	store.RUnlock()

	if !ok {
		return nil, fmt.Errorf("Invalid Key")
	}

	return value, nil
}

// Inserts(or Force Put) will insert into the Database
// or override values in the Database
func (s *Storage) Insert(key string, value interface{}) error {
	var mtx sync.Mutex

	store := s.getShard(key)
	mtx.Lock()
	defer mtx.Unlock()
	store.m[key] = value

	return nil

}

// delete removes the entry from the map completely.
// if the key doesnt exist then nothing happens
func (s *Storage) Delete(key string) {

	store := s.getShard(key)
	store.Lock()
	defer store.Unlock()
	delete(store.m, key)

	return

}

func (s *Storage) Exist(key string, value interface{}) bool {
	store := s.getShard(key)
	store.RLock()
	_, ok := store.m[key]
	store.RUnlock()
	if !ok {
		return false
	}
	return true
}
