package bolt

type meta struct {
	magic          int32
	version        int32
	mapSize        int
	free           *bucket
	main           *bucket
	lastPageNumber int
	transactionID  int
}
