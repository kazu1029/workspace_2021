package main

import (
	"fmt"
	"math/rand"
	"time"
)

// selectionSort sorts int slice by iterating and finding minimum value, then exchanging it and the first value.
// inputted slice is sorted.
func selectionSort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

func main() {
	slice := generateIntSlice(20)
	fmt.Println("\n--- Unsorted ---\n\n", slice)
	selectionSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)
}

func generateIntSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
