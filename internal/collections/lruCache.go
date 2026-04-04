package internal

import (
	"container/list"
	"errors"
	"fmt"
	"io"
)

var (
	ErrInvalidCapacity = errors.New("invalid capacity")
)

type LruCacheEntry[K comparable, V any] struct {
	k K
	v *V
}

type LruCache[K comparable, V any] struct {
	items map[K]*list.Element
	ll    *list.List
	cap   int
}

func NewLruCache[K comparable, V any](cap int) (*LruCache[K, V], error) {

	if cap < 1 {
		return nil, ErrInvalidCapacity
	}

	return &LruCache[K, V]{
		items: make(map[K]*list.Element),
		ll:    list.New(),
		cap:   cap,
	}, nil
}

func (c *LruCache[K, V]) Get(key K) *V {
	el, ok := c.items[key]
	if !ok {
		return nil
	}
	c.promote(el)
	return el.Value.(LruCacheEntry[K, V]).v
}

func (c *LruCache[K, V]) Set(key K, value *V) {
	if el, ok := c.items[key]; ok {
		el.Value = LruCacheEntry[K, V]{k: key, v: value}
		c.promote(el)
		return
	}
	if len(c.items) >= c.cap {
		back := c.ll.Back()
		ent := back.Value.(LruCacheEntry[K, V])
		delete(c.items, ent.k)
		c.ll.Remove(back)
	}
	el := c.ll.PushFront(LruCacheEntry[K, V]{k: key, v: value})
	c.items[key] = el
}

func (c *LruCache[K, V]) promote(item *list.Element) {
	c.ll.MoveToFront(item)
}

func (c *LruCache[K, V]) Format(f fmt.State, verb rune) {
	if c == nil {
		io.WriteString(f, "<nil>")
		return
	}
	switch verb {
	case 'v', 's':
		fmt.Fprintf(f, "LruCache{cap:%d ", c.cap)
		sep := ""
		for el := c.ll.Front(); el != nil; el = el.Next() {
			ent := el.Value.(LruCacheEntry[K, V])
			fmt.Fprintf(f, "%s%v:%v", sep, ent.k, ent.v)
			sep = ", "
		}
		io.WriteString(f, "}")
	default:
		fmt.Fprintf(f, "%%!%c(LruCache)", verb)
	}
}
