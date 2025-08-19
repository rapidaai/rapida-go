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
	"os"
	"path/filepath"

	lexatic_backend "github.com/rapidaai/rapida-go/rapida/clients/protos"
	rapida_constants "github.com/rapidaai/rapida-go/rapida/constants"
	"google.golang.org/protobuf/types/known/structpb"
)

// ContentBuilder builds Content objects.
type contentBuilder struct {
	name          string
	contentType   string
	contentFormat string
	content       []byte
	meta          *structpb.Struct
}

type ContentBuilder interface {
	WithFile(filePath string) (ContentBuilder, error)
	WithString(content string) ContentBuilder
	WithMeta(meta *structpb.Struct) ContentBuilder
	Build() *lexatic_backend.Content
}

// NewContentBuilder creates a new instance of ContentBuilder.
func NewContentBuilder() ContentBuilder {
	return &contentBuilder{}
}

// // WithContent sets the content field of the Content.
func (cb *contentBuilder) WithString(content string) ContentBuilder {
	cb.content = []byte(content)
	cb.contentType = rapida_constants.TEXT_CONTENT
	cb.contentFormat = rapida_constants.TEXT_CONTENT_FORMAT_RAW
	return cb
}

func (cb *contentBuilder) WithFile(filePath string) (ContentBuilder, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	cb.content = data
	cb.contentType = rapida_constants.MULTI_MEDIA_CONTENT
	cb.contentFormat = rapida_constants.MULTI_MEDIA_CONTENT_FORMAT_RAW
	// Set name from file base name
	cb.name = filepath.Base(filePath)

	// Create meta data
	metaMap := make(map[string]interface{})

	// Get file information
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	metaMap["file_size"] = fileInfo.Size()
	metaMap["file_name"] = fileInfo.Name()
	metaMap["file_ext"] = filepath.Ext(filePath)
	// Convert metaMap to structpb.Struct
	metaStruct, err := structpb.NewStruct(metaMap)
	if err != nil {
		return cb, err
	}
	cb.meta = metaStruct
	return cb, nil
}

// WithMeta sets the meta field of the Content.
func (cb *contentBuilder) WithMeta(meta *structpb.Struct) ContentBuilder {
	cb.meta = meta
	return cb
}

// Build constructs the Content object with the provided data.
func (cb *contentBuilder) Build() *lexatic_backend.Content {
	return &lexatic_backend.Content{
		Name:          cb.name,
		ContentType:   cb.contentType,
		ContentFormat: cb.contentFormat,
		Content:       cb.content,
		Meta:          cb.meta,
	}
}
