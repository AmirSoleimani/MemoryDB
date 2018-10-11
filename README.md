# MemoryDB
----
# Interfaces
  - NewMemDB
  - NewMemDBWithCap
  - Put
  - Get
  - Has
  - Keys
  - Delete
  - Batch
  -- Put
  -- Delete


### Installation

Just use go get.
```go
    go get github.com/AmirSoleimani/MemoryDB/memdb
```

### Sample

Just use go get.
```go
    func main() {
    	// sample
    	memDB := db.NewMemDB()
    	memDB.Put([]byte("blabla"), []byte("gogogo"))
    
    	a, _ := memDB.Get([]byte("blabla"))
    	fmt.Println(string(a))
    
    	batch := memDB.NewBatch()
    	batch.Put([]byte("1111"), []byte("aaaaaa"))
    	batch.Put([]byte("2222"), []byte("bbbbb"))
    	batch.Delete([]byte("blabla"))
    	batch.Write()
    
    	fmt.Println(memDB.Keys())
	}
```


