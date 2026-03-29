package internal

import "errors"

var (
	ErrInvalidCapacity = errors.New("invalid capacity")
)

type LruCache[T comparable, TT any] struct {
	items map[T]*Node[TT]
	head  *Node[TT]
	tail  *Node[TT]
	len   int
	cap   int
}

func NewLruCache[T comparable, TT any](cap int) (*LruCache[T, TT], error) {

	if cap < 1 {
		return nil, ErrInvalidCapacity
	}

	return &LruCache[T, TT]{
		items: make(map[T]*Node[TT]),
		cap:   cap,
	}, nil
}

func (c *LruCache[T, TT]) Get(key T) *TT {
	if item, ok := c.items[key]; ok {
		c.promote(item)
		return item.data
	}
	return nil
}

func (c *LruCache[T, TT]) Set(key T, value *TT) {

	if v, ok := c.items[key]; ok {

	}

	if c.len < c.cap {

	} else {

	}
}
