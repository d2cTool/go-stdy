package concurrency

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRun_EmptyTasks(t *testing.T) {
	t.Parallel()
	err := Run(nil, 3, 1)
	assert.NoError(t, err)
}

func TestRun_AllTasksExecutedOnSuccess(t *testing.T) {
	t.Parallel()

	const total = 50
	var ran atomic.Int32

	tasks := make([]Task, total)
	for i := range tasks {
		tasks[i] = func() error {
			ran.Add(1)
			return nil
		}
	}

	err := Run(tasks, 4, 2)
	require.NoError(t, err)
	assert.Equal(t, int32(total), ran.Load())
}

func TestRun_ErrorLimitExceeded(t *testing.T) {
	t.Parallel()

	errSentinel := errors.New("task failed")
	tasks := make([]Task, 20)
	for i := range tasks {
		tasks[i] = func() error { return errSentinel }
	}

	err := Run(tasks, 3, 2)
	require.ErrorIs(t, err, ErrErrorsLimitExceeded)
}

func TestRun_MNonPositive_IgnoresErrorLimit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		m    int
	}{
		{"m=0", 0},
		{"m=-1", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const total = 10
			var ran atomic.Int32

			tasks := make([]Task, total)
			for i := range tasks {
				tasks[i] = func() error {
					ran.Add(1)
					return errors.New("fail")
				}
			}

			err := Run(tasks, 2, tt.m)
			assert.NoError(t, err)
			assert.Equal(t, int32(total), ran.Load())
		})
	}
}

func TestRun_AtMostNPlusMTasksOnEarlyErrors(t *testing.T) {
	t.Parallel()

	const (
		n     = 3
		m     = 2
		total = 100
	)

	var ran atomic.Int32
	tasks := make([]Task, total)
	for i := range tasks {
		tasks[i] = func() error {
			ran.Add(1)
			return errors.New("fail")
		}
	}

	err := Run(tasks, n, m)
	require.ErrorIs(t, err, ErrErrorsLimitExceeded)
	assert.LessOrEqual(t, ran.Load(), int32(n+m),
		"при быстрых ошибках не должно выполниться больше n+m задач")
	assert.GreaterOrEqual(t, ran.Load(), int32(m),
		"должно успеть выполниться как минимум m задач с ошибкой")
}

func TestRun_FewerTasksThanWorkers(t *testing.T) {
	t.Parallel()

	var ran atomic.Int32
	tasks := []Task{
		func() error { ran.Add(1); return nil },
		func() error { ran.Add(1); return nil },
	}

	err := Run(tasks, 10, 1)
	require.NoError(t, err)
	assert.Equal(t, int32(2), ran.Load())
}

func TestRun_ConcurrencyBoundedByN(t *testing.T) {
	t.Parallel()

	const (
		n        = 4
		numTasks = 12
	)

	var inFlight atomic.Int32
	var maxObserved atomic.Int32
	release := make(chan struct{})
	started := make(chan struct{}, numTasks)

	tasks := make([]Task, numTasks)
	for i := range tasks {
		tasks[i] = func() error {
			cur := inFlight.Add(1)
			defer inFlight.Add(-1)

			for {
				prev := maxObserved.Load()
				if cur <= prev {
					break
				}
				if maxObserved.CompareAndSwap(prev, cur) {
					break
				}
			}

			started <- struct{}{}
			<-release
			return nil
		}
	}

	done := make(chan error, 1)
	go func() {
		done <- Run(tasks, n, 0)
	}()

	for range n {
		<-started
	}
	assert.GreaterOrEqual(t, maxObserved.Load(), int32(n),
		"одновременно должны работать не менее n задач")

	close(release)
	require.NoError(t, <-done)
	assert.LessOrEqual(t, maxObserved.Load(), int32(n),
		"одновременно не должно работать больше n задач")
}

func TestRun_NoGoroutineLeak(t *testing.T) {
	t.Parallel()

	var wg sync.WaitGroup
	wg.Add(1)

	tasks := []Task{
		func() error {
			wg.Done()
			return nil
		},
	}

	err := Run(tasks, 2, 1)
	require.NoError(t, err)

	wg.Wait()
}
