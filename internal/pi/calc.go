package pi

import (
	"context"
	"sync"
)

type SeriesFn func(context.Context, int) float64

func CalcParallel(ctx context.Context, n, i int) float64 {

	var res float64
	var arr []float64
	iter := i / n

	var wg sync.WaitGroup

	if n == 1 {
		wg.Go(func() {
			res = series(ctx, iter, 1, 4, 1) + series(ctx, iter, 3, 4, -1)
		})
		wg.Wait()
	} else {
		count := max(n>>1, 1)
		perShard := iter / count
		fns := split(count)
		arr = make([]float64, len(fns))
		for i, fn := range fns {
			wg.Go(func() {
				arr[i] = fn(ctx, perShard)
			})
		}
		wg.Wait()
	}

	if n > 1 {
		for _, v := range arr {
			res += v
		}
	}

	//fmt.Println("π: ", 4*res)
	return 4 * res
}

func split(count int) []SeriesFn {
	fns := make([]SeriesFn, count*2)
	step := 4.0 * float64(count)

	for i := range count {
		ii := i
		fns[2*i] = func(ctx context.Context, iter int) float64 {
			return series(ctx, iter, 1.0+4.0*float64(ii), step, 1)
		}
		fns[2*i+1] = func(ctx context.Context, iter int) float64 {
			return series(ctx, iter, 3.0+4.0*float64(ii), step, -1)
		}
	}
	return fns
}

func series(ctx context.Context, iter int, start, step, sign float64) float64 {
	res := 0.0
	d := start
	for range iter {
		if ctx.Err() != nil {
			return res
		}
		res += 1.0 / d
		d += step
	}
	return res * sign
}
