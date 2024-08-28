/*
 *  Copyright (c) 2024. Rapida
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in
 *  all copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 *  THE SOFTWARE.
 *
 *  Author: Prashant <prashant@rapida.ai>
 *
 */

package rapida_constants

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGet tests the Get method of RapidaEnvironment
func TestGet(t *testing.T) {
	tests := []struct {
		env      RapidaEnvironment
		expected string
	}{
		{PRODUCTION, "production"},
		{DEVELOPMENT, "development"},
	}

	for _, tt := range tests {
		t.Run(string(tt.env), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.env.Get())
		})
	}
}

// TestFromStr tests the FromStr method of RapidaEnvironment
func TestFromStr(t *testing.T) {
	tests := []struct {
		label    string
		expected RapidaEnvironment
	}{
		{"production", PRODUCTION},
		{"development", DEVELOPMENT},
		{"invalid", DEVELOPMENT},
		{"", DEVELOPMENT},
	}

	// Capture logs
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil) // Reset log output
	}()

	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			result := DEVELOPMENT.FromStr(tt.label)
			assert.Equal(t, tt.expected, result)

			// Check for log messages for unsupported environments
			if tt.expected == DEVELOPMENT && tt.label != "development" {
				assert.True(t, strings.Contains(buf.String(), "The environment is not supported. Only 'production' and 'development' are allowed."))
			} else {
				assert.Empty(t, buf.String())
			}
		})
	}
}

// TestGet tests the Get method of RapidaRegion
func TestRegionGet(t *testing.T) {
	tests := []struct {
		region   RapidaRegion
		expected string
	}{
		{AP, "ap"},
		{US, "us"},
		{EU, "eu"},
		{ALL, "all"},
	}

	for _, tt := range tests {
		t.Run(string(tt.region), func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.region.Get())
		})
	}
}

// TestFromStr tests the FromStr method of RapidaRegion
func TestRegionFromStr(t *testing.T) {
	tests := []struct {
		label    string
		expected RapidaRegion
	}{
		{"ap", AP},
		{"us", US},
		{"eu", EU},
		{"all", ALL},
		{"invalid", ALL},
		{"", ALL},
	}

	// Capture logs
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil) // Reset log output
	}()

	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			result := ALL.FromStr(tt.label)
			assert.Equal(t, tt.expected, result)

			// Check for log messages for unsupported regions
			if tt.expected == ALL && tt.label != "all" {
				assert.True(t, strings.Contains(buf.String(), "The region is not supported. Supported regions are 'ap', 'us', 'eu', and 'all'."))
			} else {
				assert.Empty(t, buf.String())
			}
		})
	}
}
