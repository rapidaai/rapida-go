// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"testing"
)

func TestRapidaRegion_Get(t *testing.T) {
	tests := []struct {
		name     string
		region   RapidaRegion
		expected string
	}{
		{
			name:     "AP region",
			region:   AP,
			expected: "ap",
		},
		{
			name:     "US region",
			region:   US,
			expected: "us",
		},
		{
			name:     "EU region",
			region:   EU,
			expected: "eu",
		},
		{
			name:     "ALL region",
			region:   ALL,
			expected: "all",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.region.Get()
			if result != tt.expected {
				t.Errorf("RapidaRegion.Get() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestFromRegionStr(t *testing.T) {
	tests := []struct {
		name     string
		label    string
		expected RapidaRegion
	}{
		{
			name:     "ap lowercase",
			label:    "ap",
			expected: AP,
		},
		{
			name:     "ap uppercase",
			label:    "AP",
			expected: AP,
		},
		{
			name:     "ap mixed case",
			label:    "Ap",
			expected: AP,
		},
		{
			name:     "us lowercase",
			label:    "us",
			expected: US,
		},
		{
			name:     "us uppercase",
			label:    "US",
			expected: US,
		},
		{
			name:     "eu lowercase",
			label:    "eu",
			expected: EU,
		},
		{
			name:     "eu uppercase",
			label:    "EU",
			expected: EU,
		},
		{
			name:     "all lowercase",
			label:    "all",
			expected: ALL,
		},
		{
			name:     "all uppercase",
			label:    "ALL",
			expected: ALL,
		},
		{
			name:     "invalid string defaults to ALL",
			label:    "invalid",
			expected: ALL,
		},
		{
			name:     "empty string defaults to ALL",
			label:    "",
			expected: ALL,
		},
		{
			name:     "asia-pacific defaults to ALL",
			label:    "asia-pacific",
			expected: ALL,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromRegionStr(tt.label)
			if result != tt.expected {
				t.Errorf("FromRegionStr(%q) = %q, want %q", tt.label, result, tt.expected)
			}
		})
	}
}

func TestRapidaRegionConstants(t *testing.T) {
	if AP != RapidaRegion("ap") {
		t.Errorf("AP = %q, want %q", AP, "ap")
	}
	if US != RapidaRegion("us") {
		t.Errorf("US = %q, want %q", US, "us")
	}
	if EU != RapidaRegion("eu") {
		t.Errorf("EU = %q, want %q", EU, "eu")
	}
	if ALL != RapidaRegion("all") {
		t.Errorf("ALL = %q, want %q", ALL, "all")
	}
}
