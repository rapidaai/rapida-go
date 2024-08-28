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
)

// Metric wraps around the Common.Metric protobuf message.
type Metric struct {
	Original    *lexatic_backend.Metric
	Description string
	Name        string
	Value       string
}

// NewMetric constructs a Metric wrapper from a Common.Metric protobuf message.
func NewMetric(data *lexatic_backend.Metric) *Metric {
	return &Metric{
		Original:    data,
		Description: data.Description,
		Name:        data.Name,
		Value:       data.Value,
	}
}

// ToJSON converts the original protobuf metric to a JSON string.
func (m *Metric) ToJSON() (string, error) {
	if m.Original == nil {
		return "", errors.New("no original metric to convert to JSON")
	}
	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
		Indent:          "  ",
		UseProtoNames:   true,
	}
	jsonData, err := marshaler.Marshal(m.Original)
	if err != nil {
		log.Printf("Failed to convert metric to JSON: %v", err)
		return "", err
	}
	return string(jsonData), nil
}

// ToDict converts the original protobuf metric to a map.
func (m *Metric) ToDict() (map[string]interface{}, error) {
	jsonString, err := m.ToJSON()
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
