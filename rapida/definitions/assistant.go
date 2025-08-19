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

type AssistantDefinition interface {
	GetAssistant() uint64
	GetAssistantVersion() string
}

type assistantDefinition struct {
	assistant        uint64
	assistantVersion *string
}

func NewAssistant(assistant uint64, assistantVersion *string) AssistantDefinition {
	return &assistantDefinition{
		assistant:        assistant,
		assistantVersion: assistantVersion,
	}
}

func (ed *assistantDefinition) GetAssistant() uint64 {
	return ed.assistant
}

func (ed *assistantDefinition) GetAssistantVersion() string {
	if ed.assistantVersion == nil {
		return "latest"
	}
	return *ed.assistantVersion
}
