package internal

type LruCache[T comparable, TT any] struct {
	items map[T]*Node[TT]
	head  *Node[TT]
	len   int
	cap   int
}

func NewLruCache[T comparable, TT any](cap int) *LruCache[T, TT] {
	return &LruCache[T, TT]{
		items: make(map[T]*Node[TT]),
		cap:   cap,
	}
}

func (c *LruCache[T, TT]) Get(key T) *TT {

}

func (c *LruCache[T, TT]) Set(key T, value *TT) {

	if v, ok := c.items[key]; ok {

	}

	if c.len < c.cap {

	} else {

	}
}
