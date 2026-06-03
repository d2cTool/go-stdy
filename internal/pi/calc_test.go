package pi

import (
	"context"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func leibnizPi(ctx context.Context, iter int) float64 {
	return 4 * (series(ctx, iter, 1, 4, 1) + series(ctx, iter, 3, 4, -1))
}

func TestCalcParallel_N1_MatchesSerial(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		i    int
	}{
		{"one term per sign", 1},
		{"few terms", 6},
		{"more terms", 10_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := leibnizPi(ctx, tt.i)
			got := CalcParallel(ctx, 1, tt.i)
			assert.InDelta(t, want, got, 1e-12)
		})
	}
}

func TestCalcParallel_N1_ApproximatesPi(t *testing.T) {
	ctx := context.Background()
	got := CalcParallel(ctx, 1, 1_000_000)

	assert.InDelta(t, math.Pi, got, 1e-5)
}

func TestCalcParallel_N1_KnownValue(t *testing.T) {
	ctx := context.Background()
	// 4·(1 − 1/3) = 8/3
	got := CalcParallel(ctx, 1, 1)
	assert.InDelta(t, 8.0/3.0, got, 1e-12)
}

func TestCalcParallel_ParallelMatchesScaledSerial(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		n    int
		i    int
	}{
		{"n=2", 2, 100_000},
		{"n=4", 4, 100_000},
		{"n=8", 8, 80_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, 0, tt.i%tt.n, "i должно делиться на n")
			iter := tt.i / tt.n
			want := leibnizPi(ctx, iter)
			got := CalcParallel(ctx, tt.n, tt.i)
			assert.InDelta(t, want, got, 1e-10)
		})
	}
}

func TestCalcParallel_N2_EqualsN1WithHalfIterations(t *testing.T) {
	ctx := context.Background()
	const i = 200_000

	got := CalcParallel(ctx, 2, i)
	want := CalcParallel(ctx, 1, i/2)

	assert.InDelta(t, want, got, 1e-10)
}

func TestCalcParallel_ContextCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	got := CalcParallel(ctx, 1, 1_000_000)

	assert.False(t, math.IsNaN(got))
	assert.False(t, math.IsInf(got, 0))
}

func TestCalcParallel_NoPanicForParallelModes(t *testing.T) {
	ctx := context.Background()

	for _, n := range []int{2, 4, 8} {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			require.NotPanics(t, func() {
				_ = CalcParallel(ctx, n, 1000)
			})
		})
	}
}
