package db

type kv struct {
	k, v []byte
	del  bool
}

type memBatch struct {
	db     *MemDB
	writes []kv
	size   int
}

func (b *memBatch) Put(key, value []byte) error {
	b.writes = append(b.writes, kv{key, value, false})
	b.size += len(value)
	return nil
}

func (b *memBatch) Delete(key []byte) error {
	b.writes = append(b.writes, kv{key, nil, true})
	b.size++
	return nil
}

func (b *memBatch) Write() error {
	b.db.lock.Lock()
	defer b.db.lock.Unlock()

	for _, kv := range b.writes {
		if kv.del {
			delete(b.db.db, string(kv.k))
			continue
		}
		b.db.db[string(kv.k)] = kv.v
	}
	return nil
}

func (b *memBatch) ValueSize() int {
	return b.size
}

func (b *memBatch) Reset() {
	b.writes = b.writes[:0]
	b.size = 0
}
