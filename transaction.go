package bolt

type Transaction interface {
}

type transaction struct {
	id             int
	flags          int
	db             *db
	parent         *transaction
	child          *transaction
	nextPageNumber int
	freePages      []int
	spillPages     []int
	dirtyList      []int
	reader         *reader
	// TODO: bucketxs []*bucketx
	buckets     []*bucket
	bucketFlags []int
	cursors     []*cursor
	// Implicit from slices? TODO: MDB_dbi mt_numdbs;
	mt_dirty_room int
}
