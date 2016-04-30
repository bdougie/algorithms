package main

import (
	"fmt"
	"math/rand"
)

func quickSort(items []string) []string {
	if len(items) < 2 {
		return items
	}

	left, right := 0, len(items)-1

	pivotIndex := rand.Int() % len(items)

	// Move the pivot to the right
	items = swap(items, pivotIndex, right)

	// Pile elements smaller than the pivot on the left
	for i := range items {
		if items[i] < items[right] {
			items = swap(items, i, left)
			left++
		}
	}

	// Place the pivot after the last smaller element
	items = swap(items, left, right)

	// recurse through and sort
	quickSort(items[:left])
	quickSort(items[left+1:])

	return items
}

func swap(a []string, one int, two int) []string {
	a[one], a[two] = a[two], a[one]
	return a
}

func main() {
	list := []string{"C", "B", "F", "E", "G", "A", "D"}
	fmt.Println(quickSort(list))
}
