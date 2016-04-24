package main

import (
	"fmt"
)

func fib(n int) int {
	base := []int{0, 1}
	for _, t := range base {
		if n == t {
			return n
		}
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	i := 0
	for i < 10 {
		fmt.Println(fib(i))
		i += 1
	}
}
