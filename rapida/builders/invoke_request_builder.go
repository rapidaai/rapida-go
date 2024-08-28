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

	"github.com/golang/protobuf/ptypes/any"
	rapida_definitions "github.com/rapidaai/rapida-sdk/rapida/definitions"
)

// Builder pattern for Invoke function parameters
type invokeRequestBuilder struct {
	endpoint rapida_definitions.EndpointDefinition
	ctx      context.Context
	inputs   map[string]*any.Any
	metadata map[string]*any.Any
	options  map[string]*any.Any
}

type RequestBuilder interface {
	WithContext(context.Context) RequestBuilder
	WithInputs(inputs map[string]*any.Any) RequestBuilder
	WithMetadata(metadata map[string]*any.Any) RequestBuilder
	WithOptions(options map[string]*any.Any) RequestBuilder
	Build() (context.Context, rapida_definitions.EndpointDefinition, map[string]*any.Any, map[string]*any.Any, map[string]*any.Any)
}

func InvokeRequestBuilder(endpoint rapida_definitions.EndpointDefinition) RequestBuilder {
	return &invokeRequestBuilder{
		endpoint: endpoint,
		ctx:      context.Background(),
		inputs:   make(map[string]*any.Any),
		metadata: make(map[string]*any.Any),
		options:  make(map[string]*any.Any),
	}
}

// WithInputs sets the inputs parameter
func (b *invokeRequestBuilder) WithContext(ctx context.Context) RequestBuilder {
	b.ctx = ctx
	return b
}

// WithInputs sets the inputs parameter
func (b *invokeRequestBuilder) WithInputs(inputs map[string]*any.Any) RequestBuilder {
	b.inputs = inputs
	return b
}

// WithMetadata sets the metadata parameter
func (b *invokeRequestBuilder) WithMetadata(metadata map[string]*any.Any) RequestBuilder {
	b.metadata = metadata
	return b
}

// WithOptions sets the options parameter
func (b *invokeRequestBuilder) WithOptions(options map[string]*any.Any) RequestBuilder {
	b.options = options
	return b
}

// Build returns the parameters needed for the Invoke function
func (b *invokeRequestBuilder) Build() (context.Context, rapida_definitions.EndpointDefinition, map[string]*any.Any, map[string]*any.Any, map[string]*any.Any) {
	return b.ctx, b.endpoint, b.inputs, b.metadata, b.options
}
