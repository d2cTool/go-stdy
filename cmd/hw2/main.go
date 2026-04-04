package main

import (
	"fmt"

	collections "go-stdy/internal/collections"
)

func main() {
	set := collections.NewSet[int]()
	set.Items[1] = struct{}{}
	set.Items[2] = struct{}{}
	set.Items[3] = struct{}{}

	fmt.Println(set)
}
