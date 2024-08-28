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

package rapida

import (
	"context"
	"errors"

	rapida_definitions "github.com/rapidaai/rapida-sdk/rapida/definitions"
	"google.golang.org/protobuf/types/known/anypb"
)

// RapidaClient represents a client for interacting with the Rapida API.
type rapidaClient struct {
	rapidaBridge RapidaBridge
}

type Client interface {
	Invoke(
		ctx context.Context,
		endpoint rapida_definitions.EndpointDefinition,
		inputs map[string]*anypb.Any,
		metadata map[string]*anypb.Any,
		options map[string]*anypb.Any,
	) (*rapida_definitions.InvokeResponseWrapper, error)
}

// NewRapidaClient creates a new instance of RapidaClient with the provided options.
func GetClient(options *RapidaClientOption) (*rapidaClient, error) {
	if options == nil {
		options = NewRapidaClientOption()
	}

	if options.GetRapidaApiKey() == nil || *options.GetRapidaApiKey() == "" {
		return nil, errors.New("the provided api key is invalid")
	}

	if options.GetRapidaEndpointUrl() == nil || *options.GetRapidaEndpointUrl() == "" {
		return nil, errors.New("the endpoint url is invalid")
	}

	bridge, err := NewRapidaBridge(options)
	if err != nil {
		return nil, errors.New("unable to initialize the rapida client")
	}
	return &rapidaClient{
		rapidaBridge: bridge,
	}, nil
}

// Invoke calls the Rapida API with the specified parameters.
func (client *rapidaClient) Invoke(
	ctx context.Context,
	endpoint rapida_definitions.EndpointDefinition,
	inputs map[string]*anypb.Any,
	metadata map[string]*anypb.Any,
	options map[string]*anypb.Any,
) (*rapida_definitions.InvokeResponseWrapper, error) {
	return client.rapidaBridge.InvokeWithContext(ctx, endpoint, inputs, metadata, options)
}
