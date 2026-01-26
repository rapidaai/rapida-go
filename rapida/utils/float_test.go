// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"math"
	"testing"
)

func TestAverageFloat32(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float32
		expected float32
	}{
		{
			name:     "empty slice returns zero",
			numbers:  []float32{},
			expected: 0,
		},
		{
			name:     "single element",
			numbers:  []float32{5.0},
			expected: 5.0,
		},
		{
			name:     "two elements",
			numbers:  []float32{2.0, 4.0},
			expected: 3.0,
		},
		{
			name:     "multiple elements",
			numbers:  []float32{1.0, 2.0, 3.0, 4.0, 5.0},
			expected: 3.0,
		},
		{
			name:     "negative numbers",
			numbers:  []float32{-1.0, -2.0, -3.0},
			expected: -2.0,
		},
		{
			name:     "mixed positive and negative",
			numbers:  []float32{-5.0, 0.0, 5.0},
			expected: 0.0,
		},
		{
			name:     "decimal numbers",
			numbers:  []float32{1.5, 2.5, 3.5},
			expected: 2.5,
		},
		{
			name:     "all zeros",
			numbers:  []float32{0.0, 0.0, 0.0},
			expected: 0.0,
		},
		{
			name:     "large numbers",
			numbers:  []float32{1000000.0, 2000000.0, 3000000.0},
			expected: 2000000.0,
		},
		{
			name:     "small numbers",
			numbers:  []float32{0.001, 0.002, 0.003},
			expected: 0.002,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AverageFloat32(tt.numbers)
			// Use approximate comparison for floating point
			if math.Abs(float64(result-tt.expected)) > 0.0001 {
				t.Errorf("AverageFloat32() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAverageFloat32_Precision(t *testing.T) {
	// Test that the average is calculated with reasonable precision
	numbers := []float32{0.1, 0.2, 0.3}
	result := AverageFloat32(numbers)
	expected := float32(0.2)

	// Allow for small floating point errors
	if math.Abs(float64(result-expected)) > 0.0001 {
		t.Errorf("AverageFloat32() precision issue: got %v, want approximately %v", result, expected)
	}
}
