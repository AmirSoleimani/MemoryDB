package db

//IdealBatchSize ideal batch size
const IdealBatchSize = 100 * 1024

//Putter interface
type Putter interface {
	Put(key []byte, value []byte) error
}

//Deleter interface
type Deleter interface {
	Delete(key []byte) error
}

//Batch interface
type Batch interface {
	Putter
	Deleter
	ValueSize() int // amount of data in the batch
	Write() error
	// Reset resets the batch for reuse
	Reset()
}
