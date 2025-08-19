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
	"context"
	"os"

	"github.com/golang/protobuf/ptypes/any"
	rapida_definitions "github.com/rapidaai/rapida-go/rapida/definitions"
	rapida_utils "github.com/rapidaai/rapida-go/rapida/utils"
)

// Builder pattern for Invoke function parameters
type invokeRequestBuilder struct {
	endpoint rapida_definitions.EndpointDefinition
	ctx      context.Context
	inputs   map[string]*any.Any
	metadata map[string]*any.Any
	options  map[string]*any.Any
}

type InvokeRequestBuilder interface {
	WithContext(context.Context) InvokeRequestBuilder
	WithInputs(inputs map[string]*any.Any) InvokeRequestBuilder
	WithMetadata(metadata map[string]*any.Any) InvokeRequestBuilder
	AddStringMetadata(k string, v string) error

	WithOptions(options map[string]*any.Any) InvokeRequestBuilder
	AddStringOption(k string, v string) error
	AddNumberOption(k string, v float64) error
	AddURLOption(k string, v string) error
	AddJSONOption(k string, jsonData map[string]interface{}) error
	AddByteOption(k string, audioFile []byte) error
	AddFileOption(k string, filePath string) error
	AddBooleanOption(k string, bl bool) error

	//
	AddStringInput(k string, v string) error
	AddNumberInput(k string, v float64) error
	AddURLInput(k string, v string) error
	AddJSONInput(k string, jsonData map[string]interface{}) error
	AddByteInput(k string, audioFile []byte) error
	AddFileInput(k string, filePath string) error
	AddBooleanInput(k string, bl bool) error

	Build() (context.Context, rapida_definitions.EndpointDefinition, map[string]*any.Any, map[string]*any.Any, map[string]*any.Any)
}

func NewInvokeRequestBuilder(endpoint rapida_definitions.EndpointDefinition) InvokeRequestBuilder {
	return &invokeRequestBuilder{
		endpoint: endpoint,
		ctx:      context.Background(),
		inputs:   make(map[string]*any.Any),
		metadata: make(map[string]*any.Any),
		options:  make(map[string]*any.Any),
	}
}

// WithInputs sets the inputs parameter
func (b *invokeRequestBuilder) WithContext(ctx context.Context) InvokeRequestBuilder {
	b.ctx = ctx
	return b
}

// WithInputs sets the inputs parameter
func (b *invokeRequestBuilder) WithInputs(inputs map[string]*any.Any) InvokeRequestBuilder {
	b.inputs = inputs
	return b
}

// WithStringInput adds a string input parameter to the builder
func (b *invokeRequestBuilder) AddStringInput(k string, v string) error {
	val, err := rapida_utils.StringToAny(v)
	if err != nil {
		return err
	}
	b.inputs[k] = val
	return err
}

// WithNumberInput adds a number input parameter to the builder
func (b *invokeRequestBuilder) AddNumberInput(k string, v float64) error {
	val, err := rapida_utils.Float64ToAny(v)
	if err != nil {
		return err
	}
	b.inputs[k] = val
	return err
}

// AddURLInput adds a URL input parameter to the builder
func (b *invokeRequestBuilder) AddURLInput(k string, v string) error {
	val, err := rapida_utils.StringToAny(v)
	if err != nil {
		return err
	}
	b.inputs[k] = val
	return err
}

// WithJSONInput adds a JSON input parameter to the builder
func (b *invokeRequestBuilder) AddJSONInput(k string, jsonData map[string]interface{}) error {
	val, err := rapida_utils.JSONToAny(jsonData)
	if err != nil {
		return err
	}
	b.inputs[k] = val
	return err
}

// WithAudioFileInput adds an audio file input parameter to the builder
func (b *invokeRequestBuilder) AddByteInput(k string, audioFile []byte) error {
	val, err := rapida_utils.BytesToAny(audioFile)
	if err != nil {
		return err
	}
	b.inputs[k] = val
	return err
}

func (b *invokeRequestBuilder) AddBooleanInput(k string, bl bool) error {
	val, err := rapida_utils.BoolToAny(bl)
	if err != nil {
		return err
	}
	b.inputs[k] = val
	return err
}

// WithImageFileInput adds an image file input parameter to the builder
func (b *invokeRequestBuilder) AddFileInput(k string, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return b.AddByteInput(k, data)
}

// WithMetadata sets the metadata parameter
func (b *invokeRequestBuilder) WithMetadata(metadata map[string]*any.Any) InvokeRequestBuilder {
	b.metadata = metadata
	return b
}

// WithStringInput adds a string input parameter to the builder
func (b *invokeRequestBuilder) AddStringMetadata(k string, v string) error {
	val, err := rapida_utils.StringToAny(v)
	if err != nil {
		return err
	}
	b.metadata[k] = val
	return err
}

// WithOptions sets the options parameter
func (b *invokeRequestBuilder) WithOptions(options map[string]*any.Any) InvokeRequestBuilder {
	b.options = options
	return b
}

// WithStringInput adds a string input parameter to the builder
func (b *invokeRequestBuilder) AddStringOption(k string, v string) error {
	val, err := rapida_utils.StringToAny(v)
	if err != nil {
		return err
	}
	b.options[k] = val
	return err
}

// WithNumberInput adds a number input parameter to the builder
func (b *invokeRequestBuilder) AddNumberOption(k string, v float64) error {
	val, err := rapida_utils.Float64ToAny(v)
	if err != nil {
		return err
	}
	b.options[k] = val
	return err
}

// AddURLInput adds a URL input parameter to the builder
func (b *invokeRequestBuilder) AddURLOption(k string, v string) error {
	val, err := rapida_utils.StringToAny(v)
	if err != nil {
		return err
	}
	b.options[k] = val
	return err
}

// WithJSONInput adds a JSON input parameter to the builder
func (b *invokeRequestBuilder) AddJSONOption(k string, jsonData map[string]interface{}) error {
	val, err := rapida_utils.JSONToAny(jsonData)
	if err != nil {
		return err
	}
	b.options[k] = val
	return err
}

// WithAudioFileInput adds an audio file input parameter to the builder
func (b *invokeRequestBuilder) AddByteOption(k string, audioFile []byte) error {
	val, err := rapida_utils.BytesToAny(audioFile)
	if err != nil {
		return err
	}
	b.options[k] = val
	return err
}

func (b *invokeRequestBuilder) AddBooleanOption(k string, bl bool) error {
	val, err := rapida_utils.BoolToAny(bl)
	if err != nil {
		return err
	}
	b.options[k] = val
	return err
}

// WithImageFileInput adds an image file input parameter to the builder
func (b *invokeRequestBuilder) AddFileOption(k string, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return b.AddByteOption(k, data)
}

// Build returns the parameters needed for the Invoke function
func (b *invokeRequestBuilder) Build() (context.Context, rapida_definitions.EndpointDefinition, map[string]*any.Any, map[string]*any.Any, map[string]*any.Any) {
	return b.ctx, b.endpoint, b.inputs, b.metadata, b.options
}
