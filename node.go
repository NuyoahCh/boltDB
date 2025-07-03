package bolt

const (
	bigNode = 0x01
	subNode = 0x02
	dupNode = 0x04
)

type node struct {
	lo      int
	hi      int
	flags   int
	keySize int
	data    []byte
}

func (n *node) setFlags(f int) {
}

func (n *node) size() int {
	return 0
}
