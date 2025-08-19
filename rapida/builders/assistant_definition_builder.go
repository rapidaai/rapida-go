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

type assistantBuilder struct {
	assistant        uint64
	assistantVersion *string
}

type AssistantDefinitionBuilder interface {
	WithAssistantId(assistantId uint64) AssistantDefinitionBuilder
	WithAssistant(assistant string) AssistantDefinitionBuilder
	WithAssistantVersion(version string) AssistantDefinitionBuilder
	Build() (rapida_definitions.AssistantDefinition, error)
}

// NewAssistantBuilder initializes a new builder with default values.
func NewAssistantDefinitionBuilder() AssistantDefinitionBuilder {
	return &assistantBuilder{}
}

func (b *assistantBuilder) WithAssistantId(assistantId uint64) AssistantDefinitionBuilder {
	b.assistant = assistantId
	return b
}

// WithAssistant sets the assistant value.
func (b *assistantBuilder) WithAssistant(assistant string) AssistantDefinitionBuilder {
	var err error
	b.assistant, err = strconv.ParseUint(assistant, 10, 64)
	if err != nil {
		b.assistant = 0 // Default or invalid state
	}
	return b
}

// WithAssistantVersion sets the assistant version value.
func (b *assistantBuilder) WithAssistantVersion(version string) AssistantDefinitionBuilder {
	b.assistantVersion = &version
	return b
}

// Build constructs the AssistantDefinition.
func (b *assistantBuilder) Build() (rapida_definitions.AssistantDefinition, error) {
	if b.assistant == 0 {
		return nil, errors.New("please provide an assistant to call")
	}
	return rapida_definitions.NewAssistant(b.assistant, b.assistantVersion), nil
}
