// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"testing"
)

func TestMaxUint64(t *testing.T) {
	tests := []struct {
		name     string
		a        uint64
		b        uint64
		expected uint64
	}{
		{
			name:     "a greater than b",
			a:        10,
			b:        5,
			expected: 10,
		},
		{
			name:     "b greater than a",
			a:        5,
			b:        10,
			expected: 10,
		},
		{
			name:     "equal values",
			a:        7,
			b:        7,
			expected: 7,
		},
		{
			name:     "zero and positive",
			a:        0,
			b:        5,
			expected: 5,
		},
		{
			name:     "both zeros",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "max uint64 value",
			a:        18446744073709551615,
			b:        18446744073709551614,
			expected: 18446744073709551615,
		},
		{
			name:     "large numbers",
			a:        1000000000000,
			b:        999999999999,
			expected: 1000000000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxUint64(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("MaxUint64(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMinUint64(t *testing.T) {
	tests := []struct {
		name     string
		a        uint64
		b        uint64
		expected uint64
	}{
		{
			name:     "a less than b",
			a:        5,
			b:        10,
			expected: 5,
		},
		{
			name:     "b less than a",
			a:        10,
			b:        5,
			expected: 5,
		},
		{
			name:     "equal values",
			a:        7,
			b:        7,
			expected: 7,
		},
		{
			name:     "zero and positive",
			a:        0,
			b:        5,
			expected: 0,
		},
		{
			name:     "both zeros",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "max and near-max uint64",
			a:        18446744073709551615,
			b:        18446744073709551614,
			expected: 18446744073709551614,
		},
		{
			name:     "large numbers",
			a:        1000000000000,
			b:        999999999999,
			expected: 999999999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinUint64(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("MinUint64(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkMaxUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxUint64(1000000, 999999)
	}
}

func BenchmarkMinUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinUint64(1000000, 999999)
	}
}
