package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	n := flag.Int("n", 1, "n = 1 one thread, n = 2 two threads")
	iter := flag.Int("iter", 1_000_000, "iterations")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var res float64
	var wg sync.WaitGroup

	if *n == 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res = Series(ctx, *iter, 1, 4) - Series(ctx, *iter, 3, 4)
		}()
		wg.Wait()
	} else {
		var pos, neg float64
		wg.Add(2)
		go func() {
			defer wg.Done()
			pos = Series(ctx, *iter, 1, 4)
		}()
		go func() {
			defer wg.Done()
			neg = Series(ctx, *iter, 3, 4)
		}()
		wg.Wait()
		res = pos - neg
	}

	fmt.Println("π: ", 4*res)
}

func Series(ctx context.Context, terms int, startDenom, denomStep float64) float64 {
	res := 0.0
	d := startDenom
	for range terms {
		if ctx.Err() != nil {
			return res
		}
		res += 1.0 / d
		d += denomStep
	}
	return res
}
