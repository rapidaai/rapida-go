// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"encoding/base64"
	"reflect"
	"testing"
)

func TestMergeMaps(t *testing.T) {
	tests := []struct {
		name     string
		maps     []map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name:     "empty maps",
			maps:     []map[string]interface{}{},
			expected: map[string]interface{}{},
		},
		{
			name: "single map",
			maps: []map[string]interface{}{
				{"key1": "value1", "key2": "value2"},
			},
			expected: map[string]interface{}{"key1": "value1", "key2": "value2"},
		},
		{
			name: "two maps with no overlap",
			maps: []map[string]interface{}{
				{"key1": "value1"},
				{"key2": "value2"},
			},
			expected: map[string]interface{}{"key1": "value1", "key2": "value2"},
		},
		{
			name: "two maps with overlap - later overwrites",
			maps: []map[string]interface{}{
				{"key1": "value1"},
				{"key1": "value2"},
			},
			expected: map[string]interface{}{"key1": "value2"},
		},
		{
			name: "nested maps merge",
			maps: []map[string]interface{}{
				{"nested": map[string]interface{}{"a": 1}},
				{"nested": map[string]interface{}{"b": 2}},
			},
			expected: map[string]interface{}{"nested": map[string]interface{}{"a": 1, "b": 2}},
		},
		{
			name: "deeply nested maps merge",
			maps: []map[string]interface{}{
				{"level1": map[string]interface{}{"level2": map[string]interface{}{"a": 1}}},
				{"level1": map[string]interface{}{"level2": map[string]interface{}{"b": 2}}},
			},
			expected: map[string]interface{}{"level1": map[string]interface{}{"level2": map[string]interface{}{"a": 1, "b": 2}}},
		},
		{
			name: "three maps",
			maps: []map[string]interface{}{
				{"a": 1},
				{"b": 2},
				{"c": 3},
			},
			expected: map[string]interface{}{"a": 1, "b": 2, "c": 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeMaps(tt.maps...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeMaps() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGetCaseInsensitiveKeyValue(t *testing.T) {
	tests := []struct {
		name       string
		cfg        map[string]string
		key        string
		expected   string
		expectedOk bool
	}{
		{
			name:       "exact key match",
			cfg:        map[string]string{"key1": "value1"},
			key:        "key1",
			expected:   "value1",
			expectedOk: true,
		},
		{
			name:       "uppercase key match",
			cfg:        map[string]string{"KEY1": "value1"},
			key:        "key1",
			expected:   "value1",
			expectedOk: true,
		},
		{
			name:       "key not found",
			cfg:        map[string]string{"key1": "value1"},
			key:        "key2",
			expected:   "",
			expectedOk: false,
		},
		{
			name:       "empty map",
			cfg:        map[string]string{},
			key:        "key1",
			expected:   "",
			expectedOk: false,
		},
		{
			name:       "empty key",
			cfg:        map[string]string{"": "value1"},
			key:        "",
			expected:   "value1",
			expectedOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := GetCaseInsensitiveKeyValue(tt.cfg, tt.key)
			if result != tt.expected || ok != tt.expectedOk {
				t.Errorf("GetCaseInsensitiveKeyValue() = (%q, %v), want (%q, %v)", result, ok, tt.expected, tt.expectedOk)
			}
		})
	}
}

func TestEmbeddingToFloat64(t *testing.T) {
	t.Run("float32 to float64", func(t *testing.T) {
		input := []float32{1.0, 2.5, 3.0}
		expected := []float64{1.0, 2.5, 3.0}
		result := EmbeddingToFloat64(input)
		if len(result) != len(expected) {
			t.Errorf("EmbeddingToFloat64() length = %d, want %d", len(result), len(expected))
		}
		for i := range result {
			// Use approximate comparison for floating point conversion
			diff := result[i] - expected[i]
			if diff < -0.001 || diff > 0.001 {
				t.Errorf("EmbeddingToFloat64()[%d] = %f, want %f", i, result[i], expected[i])
			}
		}
	})

	t.Run("float64 to float64", func(t *testing.T) {
		input := []float64{1.0, 2.5, 3.7}
		expected := []float64{1.0, 2.5, 3.7}
		result := EmbeddingToFloat64(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("EmbeddingToFloat64() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []float32{}
		result := EmbeddingToFloat64(input)
		if len(result) != 0 {
			t.Errorf("EmbeddingToFloat64() length = %d, want 0", len(result))
		}
	})
}

func TestEmbeddingToFloat32(t *testing.T) {
	t.Run("float64 to float32", func(t *testing.T) {
		input := []float64{1.0, 2.5, 3.7}
		expected := []float32{1.0, 2.5, 3.7}
		result := EmbeddingToFloat32(input)
		if len(result) != len(expected) {
			t.Errorf("EmbeddingToFloat32() length = %d, want %d", len(result), len(expected))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("EmbeddingToFloat32()[%d] = %f, want %f", i, result[i], expected[i])
			}
		}
	})

	t.Run("float32 to float32", func(t *testing.T) {
		input := []float32{1.0, 2.5, 3.7}
		expected := []float32{1.0, 2.5, 3.7}
		result := EmbeddingToFloat32(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("EmbeddingToFloat32() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []float64{}
		result := EmbeddingToFloat32(input)
		if len(result) != 0 {
			t.Errorf("EmbeddingToFloat32() length = %d, want 0", len(result))
		}
	})
}

func TestFloat64SliceToByteArray(t *testing.T) {
	t.Run("valid float64 slice", func(t *testing.T) {
		input := []float64{1.0, 2.0, 3.0}
		result, err := Float64SliceToByteArray(input)
		if err != nil {
			t.Errorf("Float64SliceToByteArray() error = %v", err)
		}
		// Each float64 is 8 bytes, so 3 floats = 24 bytes
		if len(result) != 24 {
			t.Errorf("Float64SliceToByteArray() length = %d, want 24", len(result))
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []float64{}
		result, err := Float64SliceToByteArray(input)
		if err != nil {
			t.Errorf("Float64SliceToByteArray() error = %v", err)
		}
		if len(result) != 0 {
			t.Errorf("Float64SliceToByteArray() length = %d, want 0", len(result))
		}
	})
}

func TestEmbeddingToBase64(t *testing.T) {
	t.Run("valid embedding", func(t *testing.T) {
		input := []float64{1.0, 2.0}
		result := EmbeddingToBase64(input)
		if result == "" {
			t.Error("EmbeddingToBase64() returned empty string")
		}
		// Verify it's valid base64
		_, err := base64.StdEncoding.DecodeString(result)
		if err != nil {
			t.Errorf("EmbeddingToBase64() returned invalid base64: %v", err)
		}
	})

	t.Run("empty embedding", func(t *testing.T) {
		input := []float64{}
		result := EmbeddingToBase64(input)
		// Should still return valid base64 for empty input
		_, err := base64.StdEncoding.DecodeString(result)
		if err != nil {
			t.Errorf("EmbeddingToBase64() returned invalid base64 for empty input: %v", err)
		}
	})
}
