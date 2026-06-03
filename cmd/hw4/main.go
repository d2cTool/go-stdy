package main

import (
	"context"
	"flag"
	"fmt"
	"go-stdy/internal/pi"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	n := flag.Int("n", 1, "n = [1, 2, ... threads]")
	iter := flag.Int("i", 1_000_000, "iterations")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	res := pi.CalcParallel(ctx, *n, *iter)
	fmt.Println("π: ", res)
}
