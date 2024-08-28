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
	"github.com/rapidaai/rapida-go/rapida"
	rapida_constants "github.com/rapidaai/rapida-go/rapida/constants"
)

// RapidaClientOptionBuilder is a builder for RapidaClientOption.
type OptionBuilder interface {
	WithApiKey(apiKey string) OptionBuilder
	WithEndpointUrl(endpointUrl string) OptionBuilder
	WithAssistantUrl(assistantUrl string) OptionBuilder
	WithRegion(region rapida_constants.RapidaRegion) OptionBuilder
	WithEnvironment(environment rapida_constants.RapidaEnvironment) OptionBuilder
	WithSecure(secure bool) OptionBuilder
	Build() *rapida.RapidaClientOption
}

type rapidaOptionClientBuilder struct {
	opts *rapida.RapidaClientOption
}

// NewRapidaClientOptionBuilder initializes a new builder with default values.
func ClientOptionBuilder() OptionBuilder {
	return &rapidaOptionClientBuilder{
		opts: rapida.NewRapidaClientOption(),
	}
}

// WithApiKey sets the API key.
func (b *rapidaOptionClientBuilder) WithApiKey(apiKey string) OptionBuilder {
	b.opts.RapidaApiKey = &apiKey
	return b
}

// WithEndpointUrl sets the endpoint URL.
func (b *rapidaOptionClientBuilder) WithEndpointUrl(endpointUrl string) OptionBuilder {
	b.opts.RapidaEndpointUrl = &endpointUrl
	return b
}

// WithAssistantUrl sets the assistant URL.
func (b *rapidaOptionClientBuilder) WithAssistantUrl(assistantUrl string) OptionBuilder {
	b.opts.RapidaAssistantUrl = &assistantUrl
	return b
}

// WithRegion sets the Rapida region.
func (b *rapidaOptionClientBuilder) WithRegion(region rapida_constants.RapidaRegion) OptionBuilder {
	b.opts.RapidaRegion = &region
	return b
}

// WithEnvironment sets the Rapida environment.
func (b *rapidaOptionClientBuilder) WithEnvironment(environment rapida_constants.RapidaEnvironment) OptionBuilder {
	b.opts.RapidaEnvironment = &environment
	return b
}

// WithSecure sets the security flag.
func (b *rapidaOptionClientBuilder) WithSecure(secure bool) OptionBuilder {
	b.opts.IsSecure = secure
	return b
}

// Build constructs the RapidaClientOption.
func (b *rapidaOptionClientBuilder) Build() *rapida.RapidaClientOption {
	return b.opts
}
