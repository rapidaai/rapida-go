// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"encoding/json"
	"testing"
)

func TestRapidaSource_Get(t *testing.T) {
	tests := []struct {
		name     string
		source   RapidaSource
		expected string
	}{
		{
			name:     "web-plugin source",
			source:   WebPlugin,
			expected: "web-plugin",
		},
		{
			name:     "debugger source",
			source:   Debugger,
			expected: "debugger",
		},
		{
			name:     "sdk source",
			source:   SDK,
			expected: "sdk",
		},
		{
			name:     "phone-call source",
			source:   PhoneCall,
			expected: "phone-call",
		},
		{
			name:     "whatsapp source",
			source:   Whatsapp,
			expected: "whatsapp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.source.Get()
			if result != tt.expected {
				t.Errorf("RapidaSource.Get() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestFromSourceStr(t *testing.T) {
	tests := []struct {
		name     string
		label    string
		expected RapidaSource
	}{
		{
			name:     "web-plugin lowercase",
			label:    "web-plugin",
			expected: WebPlugin,
		},
		{
			name:     "web-plugin uppercase",
			label:    "WEB-PLUGIN",
			expected: WebPlugin,
		},
		{
			name:     "debugger lowercase",
			label:    "debugger",
			expected: Debugger,
		},
		{
			name:     "debugger uppercase",
			label:    "DEBUGGER",
			expected: Debugger,
		},
		{
			name:     "sdk lowercase",
			label:    "sdk",
			expected: SDK,
		},
		{
			name:     "sdk uppercase",
			label:    "SDK",
			expected: SDK,
		},
		{
			name:     "phone-call lowercase",
			label:    "phone-call",
			expected: PhoneCall,
		},
		{
			name:     "phone-call uppercase",
			label:    "PHONE-CALL",
			expected: PhoneCall,
		},
		{
			name:     "whatsapp lowercase",
			label:    "whatsapp",
			expected: Whatsapp,
		},
		{
			name:     "whatsapp uppercase",
			label:    "WHATSAPP",
			expected: Whatsapp,
		},
		{
			name:     "invalid string defaults to Debugger",
			label:    "invalid",
			expected: Debugger,
		},
		{
			name:     "empty string defaults to Debugger",
			label:    "",
			expected: Debugger,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromSourceStr(tt.label)
			if result != tt.expected {
				t.Errorf("FromSourceStr(%q) = %q, want %q", tt.label, result, tt.expected)
			}
		})
	}
}

func TestRapidaSourceConstants(t *testing.T) {
	if WebPlugin != RapidaSource("web-plugin") {
		t.Errorf("WebPlugin = %q, want %q", WebPlugin, "web-plugin")
	}
	if Debugger != RapidaSource("debugger") {
		t.Errorf("Debugger = %q, want %q", Debugger, "debugger")
	}
	if SDK != RapidaSource("sdk") {
		t.Errorf("SDK = %q, want %q", SDK, "sdk")
	}
	if PhoneCall != RapidaSource("phone-call") {
		t.Errorf("PhoneCall = %q, want %q", PhoneCall, "phone-call")
	}
	if Whatsapp != RapidaSource("whatsapp") {
		t.Errorf("Whatsapp = %q, want %q", Whatsapp, "whatsapp")
	}
}

func TestRapidaSource_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		source   RapidaSource
		expected string
	}{
		{
			name:     "web-plugin",
			source:   WebPlugin,
			expected: `"web-plugin"`,
		},
		{
			name:     "debugger",
			source:   Debugger,
			expected: `"debugger"`,
		},
		{
			name:     "sdk",
			source:   SDK,
			expected: `"sdk"`,
		},
		{
			name:     "phone-call",
			source:   PhoneCall,
			expected: `"phone-call"`,
		},
		{
			name:     "whatsapp",
			source:   Whatsapp,
			expected: `"whatsapp"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := json.Marshal(tt.source)
			if err != nil {
				t.Errorf("MarshalJSON() error = %v", err)
				return
			}
			if string(result) != tt.expected {
				t.Errorf("MarshalJSON() = %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestRapidaSource_InStruct(t *testing.T) {
	type TestStruct struct {
		Source RapidaSource `json:"source"`
	}

	ts := TestStruct{Source: WebPlugin}
	result, err := json.Marshal(ts)
	if err != nil {
		t.Errorf("Marshal struct error = %v", err)
		return
	}

	expected := `{"source":"web-plugin"}`
	if string(result) != expected {
		t.Errorf("Marshal struct = %s, want %s", result, expected)
	}
}
