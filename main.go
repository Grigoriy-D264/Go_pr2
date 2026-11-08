package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements returns positive random integers, or an empty slice for size <= 0.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	for i := range data {
		// Intn excludes its upper bound, so adding one cannot overflow int.
		data[i] = rand.Intn(int(^uint(0)>>1)) + 1
	}
	return data
}

// maximum returns the largest element, or zero for an empty slice.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// maxChunks finds the maximum using CHUNKS goroutines, or returns zero for an empty slice.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	chunkSize := len(data) / CHUNKS
	maxima := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data) // Include the remainder in the last chunk.
		}
		go func(index int, chunk []int) {
			defer wg.Done()
			// Empty chunks must not introduce zero when all input values are negative.
			maxima[index] = data[0]
			if len(chunk) > 0 {
				maxima[index] = maximum(chunk)
			}
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	started := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(started).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	started = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(started).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
}
