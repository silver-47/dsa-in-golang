package testutil

import (
	"math/rand"
	"time"
)

// GenerateRandomIntSlice creates a slice of 'size' populated with random integers.
// Useful for benchmarking sorting and searching algorithms.
func GenerateRandomIntSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(1_000_000) // Random numbers up to 1,000,000
	}
	return arr
}

// GenerateSortedIntSlice creates a slice of 'size' that is already sorted.
func GenerateSortedIntSlice(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}
	return arr
}
