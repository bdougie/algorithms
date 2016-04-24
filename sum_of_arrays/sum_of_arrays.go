package main

import (
	"fmt"
	"testing"
)

func sum_fast(array []int) bool {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum == 0
}

func TimeIt(b *testing.B) {
	a := []int{-2, 1, 1}
	sum_fast(a)
}

func main() {
	// Benchmark our function
	result := testing.Benchmark(TimeIt)

	nanoseconds := float64(result.T.Nanoseconds()) / float64(result.N)
	milliseconds := nanoseconds / 1000000.0

	fmt.Printf("%13.2f ns/op | %13.10f ms/op | %d Iterations\n", nanoseconds, milliseconds, result.N)
}
