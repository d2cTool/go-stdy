package main

import (
	"fmt"
	"log"

	collections "go-stdy/internal/collections"
)

func main() {
	cache, err := collections.NewLruCache[string, int](3)
	if err != nil {
		log.Fatal(err)
	}

	a, b, c, d := 1, 2, 3, 4
	cache.Set("x", &a)
	cache.Set("y", &b)
	cache.Set("z", &c)

	fmt.Println(cache)

	if v := cache.Get("x"); v != nil {
		fmt.Println("Get x:", *v)
	}
	if v := cache.Get("missing"); v == nil {
		fmt.Println("Get missing: nil")
	}

	cache.Set("w", &d)

	if v := cache.Get("y"); v == nil {
		fmt.Println("Get y after overflow: evicted (nil)")
	}
	if v := cache.Get("w"); v != nil {
		fmt.Println("Get w:", *v)
	}

	fmt.Println(cache)
}
