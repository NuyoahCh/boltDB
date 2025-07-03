package bolt

type Bucket interface {
}

type bucket struct {
	pad               int
	flags             int
	depth             int
	branchPageCount   int
	leafPageCount     int
	overflowPageCount int
	entryPageCount    int
	rootId            int
}
