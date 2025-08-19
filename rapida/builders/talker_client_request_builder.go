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
type talkerClientRequestBuilder struct {
	assistant rapida_definitions.AssistantDefinition
	ctx       context.Context
	metadata  map[string]*any.Any
	options   map[string]*any.Any
}

type TalkerClientRequestBuilder interface {
	WithContext(context.Context) TalkerClientRequestBuilder

	WithMetadata(metadata map[string]*any.Any) TalkerClientRequestBuilder
	AddStringMetadata(k string, v string) error

	WithOptions(options map[string]*any.Any) TalkerClientRequestBuilder
	AddStringOption(k string, v string) error
	AddNumberOption(k string, v float64) error
	AddURLOption(k string, v string) error
	AddJSONOption(k string, jsonData map[string]interface{}) error
	AddByteOption(k string, audioFile []byte) error
	AddFileOption(k string, filePath string) error
	AddBooleanOption(k string, bl bool) error

	Build() (context.Context, rapida_definitions.AssistantDefinition, map[string]*any.Any, map[string]*any.Any)
}

func NewTalkerClientRequestBuilder(assistant rapida_definitions.AssistantDefinition) TalkerClientRequestBuilder {
	return &talkerClientRequestBuilder{
		assistant: assistant,
		ctx:       context.Background(),
		metadata:  make(map[string]*any.Any),
		options:   make(map[string]*any.Any),
	}
}

// WithInputs sets the inputs parameter
func (b *talkerClientRequestBuilder) WithContext(ctx context.Context) TalkerClientRequestBuilder {
	b.ctx = ctx
	return b
}

// WithMetadata sets the metadata parameter
func (b *talkerClientRequestBuilder) WithMetadata(metadata map[string]*any.Any) TalkerClientRequestBuilder {
	b.metadata = metadata
	return b
}

// WithStringInput adds a string input parameter to the builder
func (b *talkerClientRequestBuilder) AddStringMetadata(k string, v string) error {
	val, err := rapida_utils.StringToAny(v)
	if err != nil {
		return err
	}
	b.metadata[k] = val
	return err
}

// WithOptions sets the options parameter
func (b *talkerClientRequestBuilder) WithOptions(options map[string]*any.Any) TalkerClientRequestBuilder {
	b.options = options
	return b
}

// WithStringInput adds a string input parameter to the builder
func (b *talkerClientRequestBuilder) AddStringOption(k string, v string) error {
	val, err := rapida_utils.StringToAny(v)
	if err != nil {
		return err
	}
	b.options[k] = val
	return err
}

// WithNumberInput adds a number input parameter to the builder
func (b *talkerClientRequestBuilder) AddNumberOption(k string, v float64) error {
	val, err := rapida_utils.Float64ToAny(v)
	if err != nil {
		return err
	}
	b.options[k] = val
	return err
}

// AddURLInput adds a URL input parameter to the builder
func (b *talkerClientRequestBuilder) AddURLOption(k string, v string) error {
	val, err := rapida_utils.StringToAny(v)
	if err != nil {
		return err
	}
	b.options[k] = val
	return err
}

// WithJSONInput adds a JSON input parameter to the builder
func (b *talkerClientRequestBuilder) AddJSONOption(k string, jsonData map[string]interface{}) error {
	val, err := rapida_utils.JSONToAny(jsonData)
	if err != nil {
		return err
	}
	b.options[k] = val
	return err
}

// WithAudioFileInput adds an audio file input parameter to the builder
func (b *talkerClientRequestBuilder) AddByteOption(k string, audioFile []byte) error {
	val, err := rapida_utils.BytesToAny(audioFile)
	if err != nil {
		return err
	}
	b.options[k] = val
	return err
}

func (b *talkerClientRequestBuilder) AddBooleanOption(k string, bl bool) error {
	val, err := rapida_utils.BoolToAny(bl)
	if err != nil {
		return err
	}
	b.options[k] = val
	return err
}

// WithImageFileInput adds an image file input parameter to the builder
func (b *talkerClientRequestBuilder) AddFileOption(k string, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return b.AddByteOption(k, data)
}

// Build returns the parameters needed for the Invoke function
func (b *talkerClientRequestBuilder) Build() (context.Context, rapida_definitions.AssistantDefinition, map[string]*any.Any, map[string]*any.Any) {
	return b.ctx, b.assistant, b.metadata, b.options
}
