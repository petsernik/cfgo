package main

import (
	"testing"
)

func TestLowerBoundAsc(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"empty array", []int{}, 5, 0},
		{"single element found", []int{5}, 5, 0},
		{"single element not found", []int{3}, 5, 1},
		{"multiple elements found", []int{1, 3, 5, 7, 9}, 5, 2},
		{"multiple elements not found", []int{1, 3, 5, 7, 9}, 6, 3},
		{"duplicate elements", []int{1, 2, 2, 2, 3}, 2, 1},
		{"target at beginning", []int{1, 2, 3, 4, 5}, 1, 0},
		{"target at end", []int{1, 2, 3, 4, 5}, 5, 4},
		{"target greater than all", []int{1, 2, 3, 4, 5}, 10, 5},
		{"target less than all", []int{1, 2, 3, 4, 5}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lowerBoundAsc(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("LowerBoundAsc(%v, %d) = %d, expected %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}

func TestUpperBoundAsc(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"empty array", []int{}, 5, 0},
		{"single element found", []int{5}, 5, 1},
		{"single element not found", []int{3}, 5, 1},
		{"multiple elements found", []int{1, 3, 5, 7, 9}, 5, 3},
		{"multiple elements not found", []int{1, 3, 5, 7, 9}, 6, 3},
		{"duplicate elements", []int{1, 2, 2, 2, 3}, 2, 4},
		{"target at beginning", []int{1, 2, 3, 4, 5}, 1, 1},
		{"target at end", []int{1, 2, 3, 4, 5}, 5, 5},
		{"target greater than all", []int{1, 2, 3, 4, 5}, 10, 5},
		{"target less than all", []int{1, 2, 3, 4, 5}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := upperBoundAsc(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("UpperBoundAsc(%v, %d) = %d, expected %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}

func TestLowerBoundDesc(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"empty array", []int{}, 5, 0},
		{"single element found", []int{5}, 5, 0},
		{"single element not found", []int{7}, 5, 1},
		{"descending array found", []int{9, 7, 5, 3, 1}, 5, 2},
		{"descending array not found", []int{9, 7, 5, 3, 1}, 6, 2},
		{"duplicate elements", []int{3, 2, 2, 2, 1}, 2, 1},
		{"target at beginning", []int{5, 4, 3, 2, 1}, 5, 0},
		{"target at end", []int{5, 4, 3, 2, 1}, 1, 4},
		{"target greater than all", []int{5, 4, 3, 2, 1}, 10, 0},
		{"target less than all", []int{5, 4, 3, 2, 1}, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lowerBoundDesc(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("LowerBoundDesc(%v, %d) = %d, expected %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}

func TestUpperBoundDesc(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"empty array", []int{}, 5, 0},
		{"single element found", []int{5}, 5, 1},
		{"single element not found", []int{7}, 5, 1},
		{"descending array found", []int{9, 7, 5, 3, 1}, 5, 3},
		{"descending array not found", []int{9, 7, 5, 3, 1}, 6, 2},
		{"duplicate elements", []int{3, 2, 2, 2, 1}, 2, 4},
		{"target at beginning", []int{5, 4, 3, 2, 1}, 5, 1},
		{"target at end", []int{5, 4, 3, 2, 1}, 1, 5},
		{"target greater than all", []int{5, 4, 3, 2, 1}, 10, 0},
		{"target less than all", []int{5, 4, 3, 2, 1}, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := upperBoundDesc(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("UpperBoundDesc(%v, %d) = %d, expected %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}

func TestBinpow(t *testing.T) {
	tests := []struct {
		x, n     int64
		expected int64
	}{
		{2, 0, 1},
		{2, 1, 2},
		{2, 10, 1024},
		{3, 3, 27},
		{5, 0, 1},
		{1, 100, 1},
		{10, 2, 100},
		{2, 30, 1073741824},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := binpow(tt.x, tt.n)
			if result != tt.expected {
				t.Errorf("binpow(%d, %d) = %d, expected %d", tt.x, tt.n, result, tt.expected)
			}
		})
	}
}

func TestDivUp(t *testing.T) {
	tests := []struct {
		a, b     int64
		expected int64
	}{
		{1, 2, 1},   // 1/2 = 0.5 → ceil = 1
		{2, 2, 1},   // 2/2 = 1 → ceil = 1
		{1, -2, 0},  // 1/-2 = -0.5 → ceil = 0
		{2, -2, -1}, // 2/-2 = -1 → ceil = -1
		{-1, 2, 0},  // -1/2 = -0.5 → ceil = 0
		{-2, 2, -1}, // -2/2 = -1 → ceil = -1
		{-1, -2, 1}, // -1/-2 = 0.5 → ceil = 1
		{-2, -2, 1}, // -2/-2 = 1 → ceil = 1
		{5, 2, 3},   // 5/2 = 2.5 → ceil = 3
		{-5, 2, -2}, // -5/2 = -2.5 → ceil = -2
		{5, -2, -2}, // 5/-2 = -2.5 → ceil = -2
		{-5, -2, 3}, // -5/-2 = 2.5 → ceil = 3
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := divUp(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("divUp(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivDown(t *testing.T) {
	tests := []struct {
		a, b     int64
		expected int64
	}{
		{1, 2, 0},   // 1/2 = 0.5 → floor = 0
		{2, 2, 1},   // 2/2 = 1 → floor = 1
		{1, -2, -1}, // 1/-2 = -0.5 → floor = -1
		{2, -2, -1}, // 2/-2 = -1 → floor = -1
		{-1, 2, -1}, // -1/2 = -0.5 → floor = -1
		{-2, 2, -1}, // -2/2 = -1 → floor = -1
		{-1, -2, 0}, // -1/-2 = 0.5 → floor = 0
		{-2, -2, 1}, // -2/-2 = 1 → floor = 1
		{5, 2, 2},   // 5/2 = 2.5 → floor = 2
		{-5, 2, -3}, // -5/2 = -2.5 → floor = -3
		{5, -2, -3}, // 5/-2 = -2.5 → floor = -3
		{-5, -2, 2}, // -5/-2 = 2.5 → floor = 2
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := divDown(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("divDown(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestGcd(t *testing.T) {
	tests := []struct {
		a, b     int64
		expected int64
	}{
		{35 * 9, 3 * 25 * 49, 105},
		{48, 18, 6},
		{17, 13, 1},
		{100, 25, 25},
		{0, 5, 5},
		{5, 0, 5},
		{1, 1, 1},
		{54, 24, 6},
		{101, 103, 1},
		{1000000007, 1000000009, 1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := gcd(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("gcd(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDifferentTypes(t *testing.T) {
	// Test with different types
	stringArr := []string{"apple", "banana", "cherry", "date"}
	stringTarget := "cherry"

	result := lowerBoundAsc(stringArr, stringTarget)
	if result != 2 {
		t.Errorf("LowerBoundAsc(stringArr, %s) = %d, expected 2", stringTarget, result)
	}

	floatArr := []float64{1.1, 2.2, 3.3, 4.4}
	floatTarget := 3.3

	result = upperBoundAsc(floatArr, floatTarget)
	if result != 3 {
		t.Errorf("UpperBoundAsc(floatArr, %.1f) = %d, expected 3", floatTarget, result)
	}
}

func TestEdgeCases(t *testing.T) {
	// Test with empty arrays
	emptyInt := []int{}
	if lowerBoundAsc(emptyInt, 5) != 0 {
		t.Error("LowerBoundAsc with empty array should return 0")
	}

	// Test with single element
	single := []int{42}
	if lowerBoundAsc(single, 42) != 0 {
		t.Error("LowerBoundAsc with single matching element should return 0")
	}
	if upperBoundAsc(single, 42) != 1 {
		t.Error("UpperBoundAsc with single matching element should return 1")
	}
}

func TestIsMonotone(t *testing.T) {
	tests := []struct {
		name     string
		input    []int64
		expected bool
	}{
		{
			name:     "empty slice",
			input:    []int64{},
			expected: true,
		},
		{
			name:     "single element",
			input:    []int64{42},
			expected: true,
		},
		{
			name:     "all equal",
			input:    []int64{42, 42, 42, 42},
			expected: true,
		},
		{
			name:     "strictly increasing",
			input:    []int64{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "strictly decreasing",
			input:    []int64{5, 4, 3, 2, 1},
			expected: true,
		},
		{
			name:     "non-decreasing with equals",
			input:    []int64{1, 2, 2, 3, 3, 5},
			expected: true,
		},
		{
			name:     "non-increasing with equals",
			input:    []int64{5, 5, 4, 4, 3, 3},
			expected: true,
		},
		{
			name:     "not monotone (up then down)",
			input:    []int64{1, 3, 2},
			expected: false,
		},
		{
			name:     "not monotone (down then up)",
			input:    []int64{3, 1, 2},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isMonotone(tc.input); got != tc.expected {
				t.Errorf("isMonotone(%v) = %v; want %v", tc.input, got, tc.expected)
			}
			if got := isMonotoneInt64(tc.input); got != tc.expected {
				t.Errorf("isMonotoneInt64(%v) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}
