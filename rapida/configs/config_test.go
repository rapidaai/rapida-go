// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package configs

import (
	"strings"
	"testing"
)

func TestAssistantAPIEndpoint(t *testing.T) {
	if ASSISTANT_API == "" {
		t.Error("ASSISTANT_API should not be empty")
	}
	if !strings.Contains(ASSISTANT_API, "rapida.ai") {
		t.Errorf("ASSISTANT_API should contain 'rapida.ai', got %q", ASSISTANT_API)
	}
	if !strings.Contains(ASSISTANT_API, ":443") {
		t.Errorf("ASSISTANT_API should contain port 443, got %q", ASSISTANT_API)
	}
}

func TestEndpointAPIEndpoint(t *testing.T) {
	if ENDPOINT_API == "" {
		t.Error("ENDPOINT_API should not be empty")
	}
	if !strings.Contains(ENDPOINT_API, "rapida.ai") {
		t.Errorf("ENDPOINT_API should contain 'rapida.ai', got %q", ENDPOINT_API)
	}
	if !strings.Contains(ENDPOINT_API, ":443") {
		t.Errorf("ENDPOINT_API should contain port 443, got %q", ENDPOINT_API)
	}
}

func TestWebAPIEndpoint(t *testing.T) {
	if WEB_API == "" {
		t.Error("WEB_API should not be empty")
	}
	if !strings.Contains(WEB_API, "rapida.ai") {
		t.Errorf("WEB_API should contain 'rapida.ai', got %q", WEB_API)
	}
	if !strings.Contains(WEB_API, ":443") {
		t.Errorf("WEB_API should contain port 443, got %q", WEB_API)
	}
}

func TestLocalAssistantAPIEndpoint(t *testing.T) {
	if LOCAL_ASSISTANT_API == "" {
		t.Error("LOCAL_ASSISTANT_API should not be empty")
	}
	if !strings.HasPrefix(LOCAL_ASSISTANT_API, "localhost:") {
		t.Errorf("LOCAL_ASSISTANT_API should start with 'localhost:', got %q", LOCAL_ASSISTANT_API)
	}
}

func TestLocalEndpointAPIEndpoint(t *testing.T) {
	if LOCAL_ENDPOINT_API == "" {
		t.Error("LOCAL_ENDPOINT_API should not be empty")
	}
	if !strings.HasPrefix(LOCAL_ENDPOINT_API, "localhost:") {
		t.Errorf("LOCAL_ENDPOINT_API should start with 'localhost:', got %q", LOCAL_ENDPOINT_API)
	}
}

func TestLocalWebAPIEndpoint(t *testing.T) {
	if LOCAL_WEB_API == "" {
		t.Error("LOCAL_WEB_API should not be empty")
	}
	if !strings.HasPrefix(LOCAL_WEB_API, "localhost:") {
		t.Errorf("LOCAL_WEB_API should start with 'localhost:', got %q", LOCAL_WEB_API)
	}
}

func TestEndpointValues(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		expected string
	}{
		{
			name:     "ASSISTANT_API value",
			endpoint: ASSISTANT_API,
			expected: "workflow-01.rapida.ai:443",
		},
		{
			name:     "ENDPOINT_API value",
			endpoint: ENDPOINT_API,
			expected: "endpoint-01.rapida.ai:443",
		},
		{
			name:     "WEB_API value",
			endpoint: WEB_API,
			expected: "web-01.rapida.ai:443",
		},
		{
			name:     "LOCAL_ASSISTANT_API value",
			endpoint: LOCAL_ASSISTANT_API,
			expected: "localhost:9007",
		},
		{
			name:     "LOCAL_ENDPOINT_API value",
			endpoint: LOCAL_ENDPOINT_API,
			expected: "localhost:9005",
		},
		{
			name:     "LOCAL_WEB_API value",
			endpoint: LOCAL_WEB_API,
			expected: "localhost:9001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.endpoint != tt.expected {
				t.Errorf("%s = %q, want %q", tt.name, tt.endpoint, tt.expected)
			}
		})
	}
}

func TestEndpointsAreDifferent(t *testing.T) {
	// Verify that all production endpoints are unique
	endpoints := map[string]string{
		"ASSISTANT_API": ASSISTANT_API,
		"ENDPOINT_API":  ENDPOINT_API,
		"WEB_API":       WEB_API,
	}

	seen := make(map[string]string)
	for name, endpoint := range endpoints {
		if existingName, exists := seen[endpoint]; exists {
			t.Errorf("%s and %s have the same endpoint: %q", name, existingName, endpoint)
		}
		seen[endpoint] = name
	}

	// Verify that all local endpoints are unique
	localEndpoints := map[string]string{
		"LOCAL_ASSISTANT_API": LOCAL_ASSISTANT_API,
		"LOCAL_ENDPOINT_API":  LOCAL_ENDPOINT_API,
		"LOCAL_WEB_API":       LOCAL_WEB_API,
	}

	seen = make(map[string]string)
	for name, endpoint := range localEndpoints {
		if existingName, exists := seen[endpoint]; exists {
			t.Errorf("%s and %s have the same endpoint: %q", name, existingName, endpoint)
		}
		seen[endpoint] = name
	}
}

func TestLocalPortsAreValid(t *testing.T) {
	localEndpoints := []string{
		LOCAL_ASSISTANT_API,
		LOCAL_ENDPOINT_API,
		LOCAL_WEB_API,
	}

	for _, endpoint := range localEndpoints {
		parts := strings.Split(endpoint, ":")
		if len(parts) != 2 {
			t.Errorf("Local endpoint %q should have format 'host:port'", endpoint)
		}
		// Verify it's localhost
		if parts[0] != "localhost" {
			t.Errorf("Local endpoint %q should use localhost", endpoint)
		}
	}
}
