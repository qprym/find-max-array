package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(math.MaxInt) + 1
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	largest := data[0]
	for _, value := range data[1:] {
		if value > largest {
			largest = value
		}
	}
	return largest
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	chunks := min(CHUNKS, len(data))
	chunkSize := len(data) / chunks
	maxima := make([]int, chunks)
	var wg sync.WaitGroup
	wg.Add(chunks)
	for i := 0; i < chunks; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == chunks-1 {
			end = len(data)
		}
		go func(index int, chunk []int) {
			defer wg.Done()
			maxima[index] = maximum(chunk)
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	largest := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", largest, elapsed)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)
	start = time.Now()
	largest = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", largest, elapsed)
}
