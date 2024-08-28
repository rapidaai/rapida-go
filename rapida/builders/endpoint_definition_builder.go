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

package rapida_builders

import (
	"errors"
	"strconv"

	rapida_definitions "github.com/rapidaai/rapida-go/rapida/definitions"
)

type endpointBuilder struct {
	endpoint        uint64
	endpointVersion string
}

type EndpointDefinitionBuilder interface {
}

// NewEndpointBuilder initializes a new builder with default values.
func NewEndpointDefinitionBuilder() EndpointDefinitionBuilder {
	return &endpointBuilder{}
}

func (b *endpointBuilder) WithEndpointId(endpointId uint64) EndpointDefinitionBuilder {
	b.endpoint = endpointId
	return b
}

// WithEndpoint sets the endpoint value.
func (b *endpointBuilder) WithEndpoint(endpoint string) EndpointDefinitionBuilder {
	var err error
	b.endpoint, err = strconv.ParseUint(endpoint, 10, 64)
	if err != nil {
		b.endpoint = 0 // Default or invalid state
	}
	return b
}

// WithEndpointVersion sets the endpoint version value.
func (b *endpointBuilder) WithEndpointVersion(version string) EndpointDefinitionBuilder {
	b.endpointVersion = version
	return b
}

// Build constructs the EndpointDefinition.
func (b *endpointBuilder) Build() (rapida_definitions.EndpointDefinition, error) {
	if b.endpoint == 0 {
		return nil, errors.New("endpoint must be set")
	}
	return rapida_definitions.NewEndpoint(b.endpoint, b.endpointVersion), nil
}
