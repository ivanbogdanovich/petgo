package piapp

import "context"

const chunkSize int64 = 10

type Result struct {
	Pi         float64
	Iterations int64
}

type partialResult struct {
	sum        float64
	iterations int64
}

const factor = 4

func Run(ctx context.Context, workers int) Result {
	results := make(chan partialResult, workers)

	for workerID := 0; workerID < workers; workerID++ {
		start := int64(workerID) * chunkSize
		end := start + chunkSize

		go func(start, end int64) {
			sum, iterations := sumRange(ctx, start, end)

			results <- partialResult{
				sum:        sum,
				iterations: iterations,
			}
		}(start, end)
	}

	var totalSum float64
	var totalIterations int64

	for i := 0; i < workers; i++ {
		partial := <-results

		totalSum += partial.sum
		totalIterations += partial.iterations
	}

	return Result{
		Pi:         factor * totalSum,
		Iterations: totalIterations,
	}
}

// π / 4 = 1 - 1/3 + 1/5 - 1/7 + 1/9 - ...
//
// i = 0 → 2×0+1 = 1 → +1/1
// i = 1 → 2×1+1 = 3 → -1/3
// i = 2 → 2×2+1 = 5 → +1/5
// i = 3 → 2×3+1 = 7 → -1/7
func sumRange(ctx context.Context, start, end int64) (float64, int64) {
	var sum float64
	var iterations int64

	for i := start; i < end; i++ {
		select {
		case <-ctx.Done():
			return sum, iterations
		default:
		}

		term := 1.0 / float64(2*i+1)

		if i%2 == 0 {
			sum += term
		} else {
			sum -= term
		}

		iterations++
	}

	return sum, iterations
}
