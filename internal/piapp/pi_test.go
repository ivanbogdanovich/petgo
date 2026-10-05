package piapp

import (
	"context"
	"fmt"
	"math"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		workers int
	}{
		{workers: 1},
		{workers: 2},
		{workers: 4},
		{workers: 8},
	}

	for _, tt := range tests {
		t.Run(
			fmt.Sprintf("workers_%d", tt.workers),
			func(t *testing.T) {
				result := Run(
					context.Background(),
					tt.workers,
				)

				expectedIterations := int64(tt.workers) * chunkSize

				if result.Iterations != expectedIterations {
					t.Fatalf(
						"iterations: get %d, expected %d",
						result.Iterations,
						expectedIterations,
					)
				}

				const tolerance = 0.11

				if math.Abs(result.Pi-math.Pi) > tolerance {
					t.Fatalf(
						"pi: get %.10f, expected %.10f",
						result.Pi,
						math.Pi,
					)
				}
			},
		)
	}
}

var benchmarkResult Result

func BenchmarkRun(b *testing.B) {
	for _, workers := range []int{1, 2, 4, 8} {
		b.Run(
			fmt.Sprintf("workers_%d", workers),
			func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					benchmarkResult = Run(
						context.Background(),
						workers,
					)
				}
			},
		)
	}
}
