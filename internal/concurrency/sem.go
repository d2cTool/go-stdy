package concurrency

import (
	"context"
	"sync"
	"sync/atomic"
)

// Semaphore is a weighted counting semaphore.
type Semaphore interface {
	Acquire(context.Context, int64) error
	TryAcquire(int64) bool
	Release(int64)
}

// ChanSem limits concurrency via a buffered channel (one token per unit of weight).
type ChanSem struct {
	sem chan struct{}
}

func NewChanSem(n int64) *ChanSem {
	if n < 0 {
		n = 0
	}
	return &ChanSem{
		sem: make(chan struct{}, n),
	}
}

func (s *ChanSem) Acquire(ctx context.Context, n int64) error {
	if n <= 0 {
		return nil
	}
	for i := range n {
		select {
		case s.sem <- struct{}{}:
		case <-ctx.Done():
			for range i {
				<-s.sem
			}
			return ctx.Err()
		}
	}
	return nil
}

func (s *ChanSem) TryAcquire(n int64) bool {
	if n <= 0 {
		return true
	}
	for i := range n {
		select {
		case s.sem <- struct{}{}:
		default:
			for range i {
				<-s.sem
			}
			return false
		}
	}
	return true
}

func (s *ChanSem) Release(n int64) {
	if n <= 0 {
		return
	}
	for range n {
		<-s.sem
	}
}

// MutexSem is a weighted semaphore using a mutex and condition variable.
type MutexSem struct {
	mu    sync.Mutex
	avail int64
	cond  *sync.Cond
}

func NewMutexSem(n int64) *MutexSem {
	if n < 0 {
		n = 0
	}
	s := &MutexSem{avail: n}
	s.cond = sync.NewCond(&s.mu)
	return s
}

func (s *MutexSem) Acquire(ctx context.Context, n int64) error {
	if n <= 0 {
		return nil
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for s.avail < n {
		if err := ctx.Err(); err != nil {
			return err
		}
		s.cond.Wait()
	}
	s.avail -= n
	return nil
}

func (s *MutexSem) TryAcquire(n int64) bool {
	if n <= 0 {
		return true
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.avail < n {
		return false
	}
	s.avail -= n
	return true
}

func (s *MutexSem) Release(n int64) {
	if n <= 0 {
		return
	}
	s.mu.Lock()
	s.avail += n
	s.cond.Broadcast()
	s.mu.Unlock()
}

// AtomicSem uses atomic operations for the counter and a condition variable to park waiters.
type AtomicSem struct {
	mu    sync.Mutex
	avail int64
	cond  *sync.Cond
}

func NewAtomicSem(n int64) *AtomicSem {
	if n < 0 {
		n = 0
	}
	s := &AtomicSem{avail: n}
	s.cond = sync.NewCond(&s.mu)
	return s
}

func (s *AtomicSem) Acquire(ctx context.Context, n int64) error {
	if n <= 0 {
		return nil
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	for {
		if s.tryTake(n) {
			return nil
		}
		s.mu.Lock()
		for atomic.LoadInt64(&s.avail) < n {
			if err := ctx.Err(); err != nil {
				s.mu.Unlock()
				return err
			}
			s.cond.Wait()
		}
		s.mu.Unlock()
	}
}

func (s *AtomicSem) tryTake(n int64) bool {
	for {
		v := atomic.LoadInt64(&s.avail)
		if v < n {
			return false
		}
		if atomic.CompareAndSwapInt64(&s.avail, v, v-n) {
			return true
		}
	}
}

func (s *AtomicSem) TryAcquire(n int64) bool {
	if n <= 0 {
		return true
	}
	return s.tryTake(n)
}

func (s *AtomicSem) Release(n int64) {
	if n <= 0 {
		return
	}
	atomic.AddInt64(&s.avail, n)
	s.mu.Lock()
	s.cond.Broadcast()
	s.mu.Unlock()
}
