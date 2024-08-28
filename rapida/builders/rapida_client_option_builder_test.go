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
 *  Author: Prashant Srivastav
 *
 */
package rapida_builders

import (
	"testing"

	rapida_constants "github.com/rapidaai/rapida-go/rapida/constants"
	"github.com/stretchr/testify/assert"
)

func TestClientOptionBuilder_Defaults(t *testing.T) {
	builder := ClientOptionBuilder()

	opts := builder.Build()

	assert.NotNil(t, opts)
	assert.NotNil(t, opts.RapidaApiKey)
	assert.Equal(t, "connect.rapida.ai", *opts.RapidaEndpointUrl)
	assert.Equal(t, "assistant.rapida.ai:8080", *opts.RapidaAssistantUrl)
	assert.Equal(t, rapida_constants.PRODUCTION, *opts.RapidaEnvironment)
	assert.Equal(t, rapida_constants.ALL, *opts.RapidaRegion)
	assert.True(t, opts.IsSecure)
}

func TestClientOptionBuilder_WithApiKey(t *testing.T) {
	apiKey := "test_api_key"
	builder := ClientOptionBuilder()

	builder.WithApiKey(apiKey)
	opts := builder.Build()

	assert.Equal(t, &apiKey, opts.RapidaApiKey)
}

func TestClientOptionBuilder_WithEndpointUrl(t *testing.T) {
	endpointUrl := "new.endpoint.rapida.ai"
	builder := ClientOptionBuilder()

	builder.WithEndpointUrl(endpointUrl)
	opts := builder.Build()

	assert.Equal(t, &endpointUrl, opts.RapidaEndpointUrl)
}

func TestClientOptionBuilder_WithAssistantUrl(t *testing.T) {
	assistantUrl := "new.assistant.rapida.ai:8080"
	builder := ClientOptionBuilder()

	builder.WithAssistantUrl(assistantUrl)
	opts := builder.Build()

	assert.Equal(t, &assistantUrl, opts.RapidaAssistantUrl)
}

func TestClientOptionBuilder_WithRegion(t *testing.T) {
	region := rapida_constants.US
	builder := ClientOptionBuilder()

	builder.WithRegion(region)
	opts := builder.Build()

	assert.Equal(t, &region, opts.RapidaRegion)
}

func TestClientOptionBuilder_WithEnvironment(t *testing.T) {
	environment := rapida_constants.DEVELOPMENT
	builder := ClientOptionBuilder()

	builder.WithEnvironment(environment)
	opts := builder.Build()

	assert.Equal(t, &environment, opts.RapidaEnvironment)
}

func TestClientOptionBuilder_WithSecure(t *testing.T) {
	builder := ClientOptionBuilder()

	builder.WithSecure(false)
	opts := builder.Build()

	assert.False(t, opts.IsSecure)
}

func TestClientOptionBuilder_Integration(t *testing.T) {
	apiKey := "integrated_test_key"
	endpointUrl := "integrated.endpoint.rapida.ai"
	assistantUrl := "integrated.assistant.rapida.ai:8080"
	region := rapida_constants.AP
	environment := rapida_constants.DEVELOPMENT
	isSecure := false

	builder := ClientOptionBuilder()

	opts := builder.
		WithApiKey(apiKey).
		WithEndpointUrl(endpointUrl).
		WithAssistantUrl(assistantUrl).
		WithRegion(region).
		WithEnvironment(environment).
		WithSecure(isSecure).
		Build()

	assert.Equal(t, &apiKey, opts.RapidaApiKey)
	assert.Equal(t, &endpointUrl, opts.RapidaEndpointUrl)
	assert.Equal(t, &assistantUrl, opts.RapidaAssistantUrl)
	assert.Equal(t, &region, opts.RapidaRegion)
	assert.Equal(t, &environment, opts.RapidaEnvironment)
	assert.False(t, opts.IsSecure)
}
