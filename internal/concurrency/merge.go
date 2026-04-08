package concurrency

import (
	"sync"
)

func or(channels ...<-chan any) <-chan any {
	out := make(chan any)
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, ch := range channels {
		go func(ch <-chan any) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
