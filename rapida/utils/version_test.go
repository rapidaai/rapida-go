// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"testing"
)

func TestGetVersionDefinition(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected *uint64
	}{
		{
			name:     "empty string returns nil",
			version:  "",
			expected: nil,
		},
		{
			name:     "latest returns nil",
			version:  "latest",
			expected: nil,
		},
		{
			name:     "valid version with prefix",
			version:  "vrsn_123",
			expected: ptrUint64(123),
		},
		{
			name:     "valid version without prefix",
			version:  "456",
			expected: ptrUint64(456),
		},
		{
			name:     "invalid version returns nil",
			version:  "invalid",
			expected: nil,
		},
		{
			name:     "version with prefix and invalid number",
			version:  "vrsn_abc",
			expected: nil,
		},
		{
			name:     "zero version",
			version:  "vrsn_0",
			expected: ptrUint64(0),
		},
		{
			name:     "large version number",
			version:  "vrsn_18446744073709551615",
			expected: ptrUint64(18446744073709551615),
		},
		{
			name:     "negative number string",
			version:  "vrsn_-1",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetVersionDefinition(tt.version)
			if tt.expected == nil {
				if result != nil {
					t.Errorf("GetVersionDefinition(%q) = %v, want nil", tt.version, *result)
				}
			} else {
				if result == nil {
					t.Errorf("GetVersionDefinition(%q) = nil, want %v", tt.version, *tt.expected)
				} else if *result != *tt.expected {
					t.Errorf("GetVersionDefinition(%q) = %v, want %v", tt.version, *result, *tt.expected)
				}
			}
		})
	}
}

func TestVersionPrefix(t *testing.T) {
	if VERSION_PREFIX != "vrsn_" {
		t.Errorf("VERSION_PREFIX = %q, want %q", VERSION_PREFIX, "vrsn_")
	}
}

func ptrUint64(v uint64) *uint64 {
	return &v
}
