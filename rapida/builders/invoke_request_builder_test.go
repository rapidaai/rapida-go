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
package rapdia_builders

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
)

// MockEndpointDefinition is a mock implementation of the EndpointDefinition interface
type MockEndpointDefinition struct {
	endpoint        uint64
	endpointVersion string
}

// NewMockEndpointDefinition creates a new instance of MockEndpointDefinition
func NewMockEndpointDefinition(endpoint uint64, version string) *MockEndpointDefinition {
	return &MockEndpointDefinition{
		endpoint:        endpoint,
		endpointVersion: version,
	}
}

// GetEndpoint returns the mock endpoint
func (m *MockEndpointDefinition) GetEndpoint() uint64 {
	return m.endpoint
}

// GetEndpointVersion returns the mock endpoint version
func (m *MockEndpointDefinition) GetEndpointVersion() string {
	return m.endpointVersion
}

func TestInvokeRequestBuilder(t *testing.T) {
	// Create a test endpoint using the mock
	mockEndpoint := NewMockEndpointDefinition(12345, "v1")

	// Initialize the builder
	builder := InvokeRequestBuilder(mockEndpoint)

	// Test default values
	t.Run("Test Default Values", func(t *testing.T) {
		ctx, endpoint, inputs, metadata, options := builder.Build()
		assert.Equal(t, context.Background(), ctx)
		assert.Equal(t, mockEndpoint, endpoint)
		assert.Empty(t, inputs)
		assert.Empty(t, metadata)
		assert.Empty(t, options)
	})

	// Test WithContext
	t.Run("Test WithContext", func(t *testing.T) {
		newCtx := context.WithValue(context.Background(), "key", "value")
		builder.WithContext(newCtx)
		ctx, _, _, _, _ := builder.Build()
		assert.Equal(t, newCtx, ctx)
	})

	// Test WithInputs
	t.Run("Test WithInputs", func(t *testing.T) {
		inputs := map[string]*any.Any{
			"key1": {Value: []byte("value1")},
			"key2": {Value: []byte("value2")},
		}
		builder.WithInputs(inputs)
		_, _, gotInputs, _, _ := builder.Build()
		assert.Equal(t, inputs, gotInputs)
	})

	// Test WithMetadata
	t.Run("Test WithMetadata", func(t *testing.T) {
		metadata := map[string]*any.Any{
			"key1": {Value: []byte("value1")},
		}
		builder.WithMetadata(metadata)
		_, _, _, gotMetadata, _ := builder.Build()
		assert.Equal(t, metadata, gotMetadata)
	})

	// Test WithOptions
	t.Run("Test WithOptions", func(t *testing.T) {
		options := map[string]*any.Any{
			"key1": {Value: []byte("value1")},
		}
		builder.WithOptions(options)
		_, _, _, _, gotOptions := builder.Build()
		assert.Equal(t, options, gotOptions)
	})

	// Test Method Chaining
	t.Run("Test Method Chaining", func(t *testing.T) {
		newCtx := context.WithValue(context.Background(), "key", "value")
		inputs := map[string]*any.Any{
			"inputKey": {Value: []byte("inputValue")},
		}
		metadata := map[string]*any.Any{
			"metadataKey": {Value: []byte("metadataValue")},
		}
		options := map[string]*any.Any{
			"optionsKey": {Value: []byte("optionsValue")},
		}

		builder.
			WithContext(newCtx).
			WithInputs(inputs).
			WithMetadata(metadata).
			WithOptions(options)

		ctx, endpoint, gotInputs, gotMetadata, gotOptions := builder.Build()
		assert.Equal(t, newCtx, ctx)
		assert.Equal(t, mockEndpoint, endpoint)
		assert.Equal(t, inputs, gotInputs)
		assert.Equal(t, metadata, gotMetadata)
		assert.Equal(t, options, gotOptions)
	})

	// Test Build Without Setting Any Optional Fields
	t.Run("Test Build Without Optional Fields", func(t *testing.T) {
		builder := InvokeRequestBuilder(mockEndpoint)
		ctx, endpoint, inputs, metadata, options := builder.Build()
		assert.Equal(t, context.Background(), ctx)
		assert.Equal(t, mockEndpoint, endpoint)
		assert.Empty(t, inputs)
		assert.Empty(t, metadata)
		assert.Empty(t, options)
	})
}
