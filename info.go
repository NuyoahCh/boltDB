package bolt

type info struct {
	MapSize           int
	LastPageID        int
	LastTransactionID int
	MaxReaders        int
	ReaderCount       int
}
