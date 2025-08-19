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

import lexatic_backend "github.com/rapidaai/rapida-go/rapida/clients/protos"

type MessageBuilder interface {
	WithRole(role string) MessageBuilder
	AddContent(content *lexatic_backend.Content) MessageBuilder
	Clear() MessageBuilder
	Build() *lexatic_backend.Message
}

type messageBuilder struct {
	role     string
	contents []*lexatic_backend.Content
}

// NewMessageBuilder creates a new instance of MessageBuilder.
func NewMessageBuilder() MessageBuilder {
	return &messageBuilder{
		role:     "user",
		contents: make([]*lexatic_backend.Content, 0),
	}
}

// WithRole sets the role field of the Message.
func (b *messageBuilder) WithRole(role string) MessageBuilder {
	b.role = role
	return b
}

// AddContent adds a Content to the contents list.
func (b *messageBuilder) AddContent(content *lexatic_backend.Content) MessageBuilder {
	b.contents = append(b.contents, content)
	return b
}

func (b *messageBuilder) Clear() MessageBuilder {
	b.contents = make([]*lexatic_backend.Content, 0)
	return b
}

// Build constructs the Message object with the provided data.
func (b *messageBuilder) Build() *lexatic_backend.Message {
	return &lexatic_backend.Message{
		Role:     b.role,
		Contents: b.contents,
	}
}
