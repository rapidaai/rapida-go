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

package rapida_definitions

import (
	"errors"
	"fmt"
	"strconv"
)

type EndpointDefinition interface {
	GetEndpoint() uint64
	GetEndpointVersion() string
}

type endpointDefinition struct {
	endpoint        uint64
	endpointVersion string
}

func NewEndpoint(endpoint interface{}, endpointVersion string) (EndpointDefinition, error) {
	var parsedEndpoint uint64
	switch v := endpoint.(type) {
	case uint64:
		parsedEndpoint = v
	case string:
		var err error
		parsedEndpoint, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid endpoint format: %v", err)
		}
	default:
		return nil, errors.New("unsupported endpoint type: must be uint64 or string")
	}
	return &endpointDefinition{
		endpoint:        parsedEndpoint,
		endpointVersion: endpointVersion,
	}, nil
}

func (ed *endpointDefinition) GetEndpoint() uint64 {
	return ed.endpoint
}

func (ed *endpointDefinition) GetEndpointVersion() string {
	return ed.endpointVersion
}
