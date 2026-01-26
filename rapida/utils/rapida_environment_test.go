// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"testing"
)

func TestRapidaEnvironment_Get(t *testing.T) {
	tests := []struct {
		name     string
		env      RapidaEnvironment
		expected string
	}{
		{
			name:     "production environment",
			env:      PRODUCTION,
			expected: "production",
		},
		{
			name:     "development environment",
			env:      DEVELOPMENT,
			expected: "development",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.env.Get()
			if result != tt.expected {
				t.Errorf("RapidaEnvironment.Get() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestFromEnvironmentStr(t *testing.T) {
	tests := []struct {
		name     string
		label    string
		expected RapidaEnvironment
	}{
		{
			name:     "production lowercase",
			label:    "production",
			expected: PRODUCTION,
		},
		{
			name:     "production uppercase",
			label:    "PRODUCTION",
			expected: PRODUCTION,
		},
		{
			name:     "production mixed case",
			label:    "ProDuCtIoN",
			expected: PRODUCTION,
		},
		{
			name:     "development lowercase",
			label:    "development",
			expected: DEVELOPMENT,
		},
		{
			name:     "development uppercase",
			label:    "DEVELOPMENT",
			expected: DEVELOPMENT,
		},
		{
			name:     "development mixed case",
			label:    "DeVeLoPmEnT",
			expected: DEVELOPMENT,
		},
		{
			name:     "invalid string defaults to development",
			label:    "invalid",
			expected: DEVELOPMENT,
		},
		{
			name:     "empty string defaults to development",
			label:    "",
			expected: DEVELOPMENT,
		},
		{
			name:     "staging defaults to development",
			label:    "staging",
			expected: DEVELOPMENT,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromEnvironmentStr(tt.label)
			if result != tt.expected {
				t.Errorf("FromEnvironmentStr(%q) = %q, want %q", tt.label, result, tt.expected)
			}
		})
	}
}

func TestRapidaEnvironmentConstants(t *testing.T) {
	if PRODUCTION != RapidaEnvironment("production") {
		t.Errorf("PRODUCTION = %q, want %q", PRODUCTION, "production")
	}
	if DEVELOPMENT != RapidaEnvironment("development") {
		t.Errorf("DEVELOPMENT = %q, want %q", DEVELOPMENT, "development")
	}
}
