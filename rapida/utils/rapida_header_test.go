// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"testing"
)

func TestHeaderConstants(t *testing.T) {
	tests := []struct {
		name     string
		header   string
		expected string
	}{
		{
			name:     "authorization header",
			header:   HEADER_AUTHORIZATION,
			expected: "authorization",
		},
		{
			name:     "project id header",
			header:   HEADER_PROJECT_ID,
			expected: "x-project-id",
		},
		{
			name:     "api key header",
			header:   HEADER_API_KEY,
			expected: "x-api-key",
		},
		{
			name:     "auth key header",
			header:   HEADER_AUTH_KEY,
			expected: "x-auth-id",
		},
		{
			name:     "source key header",
			header:   HEADER_SOURCE_KEY,
			expected: "x-client-source",
		},
		{
			name:     "environment key header",
			header:   HEADER_ENVIRONMENT_KEY,
			expected: "x-rapida-environment",
		},
		{
			name:     "region key header",
			header:   HEADER_REGION_KEY,
			expected: "x-rapida-region",
		},
		{
			name:     "user agent header",
			header:   HEADER_USER_AGENT,
			expected: "x-user-agent",
		},
		{
			name:     "language header",
			header:   HEADER_LANGUAGE,
			expected: "x-language",
		},
		{
			name:     "platform header",
			header:   HEADER_PLATFORM,
			expected: "x-platform",
		},
		{
			name:     "screen width header",
			header:   HEADER_SCREEN_WIDTH,
			expected: "x-screen-width",
		},
		{
			name:     "screen height header",
			header:   HEADER_SCREEN_HEIGHT,
			expected: "x-screen-height",
		},
		{
			name:     "window width header",
			header:   HEADER_WINDOW_WIDTH,
			expected: "x-window-width",
		},
		{
			name:     "window height header",
			header:   HEADER_WINDOW_HEIGHT,
			expected: "x-window-height",
		},
		{
			name:     "timezone header",
			header:   HEADER_TIMEZONE,
			expected: "x-timezone",
		},
		{
			name:     "color depth header",
			header:   HEADER_COLOR_DEPTH,
			expected: "x-color-depth",
		},
		{
			name:     "device memory header",
			header:   HEADER_DEVICE_MEMORY,
			expected: "x-device-memory",
		},
		{
			name:     "hardware concurrency header",
			header:   HEADER_HARDWARE_CONCURRENCY,
			expected: "x-hardware-concurrency",
		},
		{
			name:     "connection type header",
			header:   HEADER_CONNECTION_TYPE,
			expected: "x-connection-type",
		},
		{
			name:     "connection effective type header",
			header:   HEADER_CONNECTION_EFFECTIVE_TYPE,
			expected: "x-connection-effective-type",
		},
		{
			name:     "cookies enabled header",
			header:   HEADER_COOKIES_ENABLED,
			expected: "x-cookies-enabled",
		},
		{
			name:     "do not track header",
			header:   HEADER_DO_NOT_TRACK,
			expected: "x-do-not-track",
		},
		{
			name:     "referrer header",
			header:   HEADER_REFERRER,
			expected: "x-referrer",
		},
		{
			name:     "remote url header",
			header:   HEADER_REMOTE_URL,
			expected: "x-remote-url",
		},
		{
			name:     "latitude header",
			header:   HEADER_LATITUDE,
			expected: "x-latitude",
		},
		{
			name:     "longitude header",
			header:   HEADER_LONGITUDE,
			expected: "x-longitude",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.header != tt.expected {
				t.Errorf("Header constant = %q, want %q", tt.header, tt.expected)
			}
		})
	}
}

func TestHeaderConstants_NotEmpty(t *testing.T) {
	headers := []string{
		HEADER_AUTHORIZATION,
		HEADER_PROJECT_ID,
		HEADER_API_KEY,
		HEADER_AUTH_KEY,
		HEADER_SOURCE_KEY,
		HEADER_ENVIRONMENT_KEY,
		HEADER_REGION_KEY,
		HEADER_USER_AGENT,
		HEADER_LANGUAGE,
		HEADER_PLATFORM,
		HEADER_SCREEN_WIDTH,
		HEADER_SCREEN_HEIGHT,
		HEADER_WINDOW_WIDTH,
		HEADER_WINDOW_HEIGHT,
		HEADER_TIMEZONE,
		HEADER_COLOR_DEPTH,
		HEADER_DEVICE_MEMORY,
		HEADER_HARDWARE_CONCURRENCY,
		HEADER_CONNECTION_TYPE,
		HEADER_CONNECTION_EFFECTIVE_TYPE,
		HEADER_COOKIES_ENABLED,
		HEADER_DO_NOT_TRACK,
		HEADER_REFERRER,
		HEADER_REMOTE_URL,
		HEADER_LATITUDE,
		HEADER_LONGITUDE,
	}

	for _, header := range headers {
		if header == "" {
			t.Errorf("Header constant should not be empty")
		}
	}
}
