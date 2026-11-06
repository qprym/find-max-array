package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-5, 0, 1, 100} {
		data := generateRandomElements(size)
		if len(data) != max(0, size) {
			t.Fatalf("generateRandomElements(%d): получено элементов: %d", size, len(data))
		}
		for i, value := range data {
			if value <= 0 {
				t.Fatalf("data[%d] = %d, ожидалось положительное число", i, value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{7}, 7},
		{"negative", []int{-8, -2, -6}, -2},
		{"first", []int{9, 5, 3}, 9},
		{"repeated", []int{1, 8, 4, 8, 2, 3, 5, 6}, 8},
		{"last in remainder", []int{5, 2, 7, 1, 4, 3, 2, 6, 9}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}
