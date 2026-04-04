package main

import (
	"flag"
	"fmt"
	"os"

	"go-stdy/internal/spacker"
)

func main() {
	pack := flag.Bool("pack", false, "pack string")
	unpack := flag.Bool("unpack", false, "unpack string")
	input := flag.String("input", "", "input string")

	flag.Parse()

	if *pack == *unpack {
		fmt.Fprintln(os.Stderr, "specify exactly one of: --pack or --unpack")
		os.Exit(1)
	}

	if *pack {
		fmt.Print(spacker.Pack(*input))
		return
	}

	out, err := spacker.Unpack(*input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Print(out)
}
