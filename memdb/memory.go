package db

import (
	"errors"
	"sync"
)

//MemDB memory database
type MemDB struct {
	db   map[string][]byte
	lock sync.RWMutex
}

//NewMemDB create new memory database
func NewMemDB() *MemDB {
	return &MemDB{
		db: make(map[string][]byte),
	}
}

//NewMemDBWithCap create new memory database with capacity
func NewMemDBWithCap(size int) *MemDB {
	return &MemDB{
		db: make(map[string][]byte, size),
	}
}

//Put new data into memory db
func (db *MemDB) Put(key []byte, value []byte) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	db.db[string(key)] = value
	return nil
}

//Has check key exist in memory db
func (db *MemDB) Has(key []byte) (bool, error) {
	db.lock.RLock()
	defer db.lock.RUnlock()

	_, ok := db.db[string(key)]
	return ok, nil
}

//Get key from memory db
func (db *MemDB) Get(key []byte) ([]byte, error) {
	db.lock.RLock()
	defer db.lock.RUnlock()

	if entry, ok := db.db[string(key)]; ok {
		return entry, nil
	}
	return nil, errors.New("not found")
}

//Keys return all keys in memory db
func (db *MemDB) Keys() [][]byte {
	db.lock.RLock()
	defer db.lock.RUnlock()

	keys := [][]byte{}
	for key := range db.db {
		keys = append(keys, []byte(key))
	}
	return keys
}

//Delete key from memory db
func (db *MemDB) Delete(key []byte) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	delete(db.db, string(key))
	return nil
}

//Close close memory db
func (db *MemDB) Close() {
	//TODO: save data into file
	//...
}

//NewBatch batch W
func (db *MemDB) NewBatch() Batch {
	return &memBatch{db: db}
}

//Len return length of keys
func (db *MemDB) Len() int { return len(db.db) }
