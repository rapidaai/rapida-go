// Copyright (c) 2024. Rapida
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
//
// Author: Prashant <prashant@rapida.ai>

package rapida_definitions

import (
	"encoding/json"
	"errors"
	"log"

	lexatic_backend "github.com/rapidaai/rapida-sdk/rapida/clients/protos"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

// Content is a struct that wraps around the Common.Content protobuf message.
type Content struct {
	Original      *lexatic_backend.Content
	Content       []byte
	ContentFormat string
	ContentType   string
	Meta          *structpb.Struct
}

// NewContent constructs a Content wrapper from a Common.Content protobuf message.
func NewContent(data *lexatic_backend.Content) *Content {
	return &Content{
		Original:      data,
		Content:       data.Content,
		ContentFormat: data.ContentFormat,
		ContentType:   data.ContentType,
		Meta:          data.Meta,
	}
}

// ToText converts the content to a text string if the format and type are text-based.
func (c *Content) ToText() (string, error) {
	if c.ContentType == "text" && c.ContentFormat == "raw" {
		return string(c.Content), nil
	}
	log.Println("ToText should always get called for text format content")
	return "", errors.New("invalid content type or format for ToText")
}

// ToJSON converts the original protobuf content to a JSON string.
func (c *Content) ToJSON() (string, error) {
	if c.Original == nil {
		return "", errors.New("no original content to convert to JSON")
	}
	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
		Indent:          "  ",
		UseProtoNames:   true,
	}
	jsonData, err := marshaler.Marshal(c.Original)
	if err != nil {
		log.Printf("Failed to convert content to JSON: %v", err)
		return "", err
	}
	return string(jsonData), nil
}

// ToDict converts the original protobuf content to a map.
func (c *Content) ToDict() (map[string]interface{}, error) {
	jsonString, err := c.ToJSON()
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &result); err != nil {
		log.Printf("Failed to convert JSON to map: %v", err)
		return nil, err
	}

	return result, nil
}
