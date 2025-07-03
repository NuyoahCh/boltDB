package bolt

type Cursor interface {
	First() error
	FirstDup() error
	Get() ([]byte, []byte, error)
	GetRange() ([]byte, []byte, error)
	Current() ([]byte, []byte, error)
	Last()
	LastDup()
	Next() ([]byte, []byte, error)
	NextDup() ([]byte, []byte, error)
	NextNoDup() ([]byte, []byte, error)
	Prev() ([]byte, []byte, error)
	PrevDup() ([]byte, []byte, error)
	PrevNoDup() ([]byte, []byte, error)
	Set() ([]byte, []byte, error)
	SetRange() ([]byte, []byte, error)
}

type cursor struct {
	flags       int
	next        *cursor
	backup      *cursor
	xcursor     *xcursor
	transaction *transaction
	bucketId    int
	bucket      *bucket
	//bucketx     *bucketx
	bucketFlag int
	snum       int
	top        int
	page       []*page
	ki         []int /**< stack of page indices */
}
