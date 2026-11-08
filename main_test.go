package main

import (
	"slices"
	"strconv"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 100} {
		t.Run(strconv.Itoa(size), func(t *testing.T) {
			data := generateRandomElements(size)
			wantLen := size
			if wantLen < 0 {
				wantLen = 0
			}
			if len(data) != wantLen {
				t.Fatalf("length = %d, want %d", len(data), wantLen)
			}
			for i, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d, want a positive integer", i, value)
				}
			}
		})
	}
}

func TestMaximumFunctions(t *testing.T) {
	minInt := -int(^uint(0)>>1) - 1
	maxInt := int(^uint(0) >> 1)
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{7}, 7},
		{"zeros", []int{0, 0}, 0},
		{"negative", []int{-8, -2, -10}, -2},
		{"equal", []int{3, 3, 3}, 3},
		{"first", []int{9, 1, 2}, 9},
		{"last", []int{1, 2, 9}, 9},
		{"eight", []int{1, 2, 3, 4, 8, 6, 7, 5}, 8},
		{"remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, 99},
		{"negative remainder", []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}, -1},
		{"minimum int", []int{minInt}, minInt},
		{"maximum int", []int{0, maxInt, minInt}, maxInt},
	}
	for name, find := range map[string]func([]int) int{"maximum": maximum, "maxChunks": maxChunks} {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					before := slices.Clone(tt.data)
					if got := find(tt.data); got != tt.want {
						t.Errorf("maximum = %d, want %d", got, tt.want)
					}
					if !slices.Equal(tt.data, before) {
						t.Error("input slice was modified")
					}
				})
			}
		})
	}
}

func TestMaxChunksEveryPosition(t *testing.T) {
	// Exercise chunk boundaries and remainders, including slices shorter than CHUNKS.
	for size := 1; size <= 3*CHUNKS+1; size++ {
		for position := 0; position < size; position++ {
			data := make([]int, size)
			data[position] = 42
			if got := maxChunks(data); got != 42 {
				t.Fatalf("size %d, maximum at %d: got %d, want 42", size, position, got)
			}
		}
	}
}
