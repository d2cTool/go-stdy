package concurrency

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type WorkerPool struct {
	tasks  <-chan Task
	errors atomic.Int32
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	m      int
}

func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return nil
	}
	if n <= 0 {
		n = 1
	}

	tasksChan := make(chan Task)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool := &WorkerPool{
		tasks:  tasksChan,
		ctx:    ctx,
		cancel: cancel,
		m:      m,
	}

	for range n {
		pool.wg.Add(1)
		go pool.worker()
	}

dispatch:
	for _, task := range tasks {
		select {
		case tasksChan <- task:
		case <-pool.ctx.Done():
			break dispatch
		}
	}
	close(tasksChan)

	pool.wg.Wait()

	if m > 0 && pool.errors.Load() >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func (pool *WorkerPool) worker() {
	defer pool.wg.Done()

	for {
		select {
		case <-pool.ctx.Done():
			return
		case task, ok := <-pool.tasks:
			if !ok {
				return
			}
			if pool.ctx.Err() != nil {
				continue
			}
			pool.run(task)
		}
	}
}

func (pool *WorkerPool) run(task Task) {
	if err := task(); err != nil && pool.m > 0 {
		if pool.errors.Add(1) >= int32(pool.m) {
			pool.cancel()
		}
	}
}
