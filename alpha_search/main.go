package main

import (
	"fmt"
	"strings"
)

func LinearSearchIterations(l *[]string, s string) int {
	iterations := 0

	for _, i := range *l {
		iterations = iterations + 1

		if i == s {
			return iterations
		}
	}

	return iterations
}

func BinarySearchIterations(list *[]string, s string) int {
	low := 0
	high := len(*list) - 1
	mid := (low + high) / 2

	switch strings.Compare((*list)[mid], s) {
	case 0:
		low = mid
	case 1:
		high = mid - 1
	case -1:
		low = mid + 1
	default:
		return mid
	}

	binary := (*list)[low : high+1]
	iterations := LinearSearchIterations(&binary, s)

	return iterations
}

func main() {
	var list = []string{"A", "B", "C", "D", "E", "F", "G"}

	linearCount := LinearSearchIterations(&list, "G")
	fmt.Println("linear", linearCount)

	binaryCount := BinarySearchIterations(&list, "G")
	fmt.Println("binary", binaryCount)
}
