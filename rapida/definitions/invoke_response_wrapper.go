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

// InvokeResponseWrapper is a struct that wraps around the response from the Invoker API.
type InvokeResponseWrapper struct {
	Data    *lexatic_backend.CallerResponse
	Success bool
	Code    int32
	Error   *lexatic_backend.InvokerError
}

// NewInvokeResponseWrapper constructs an InvokeResponseWrapper from an InvokerApi.InvokeResponse.
func NewInvokeResponseWrapper(data *lexatic_backend.InvokeResponse) *InvokeResponseWrapper {
	return &InvokeResponseWrapper{
		Data:    data.GetData(),
		Success: data.GetSuccess(),
		Code:    data.GetCode(),
		Error:   data.GetError(),
	}
}

// GetTimeTaken returns the time taken for the operation.
func (irw *InvokeResponseWrapper) GetTimeTaken() uint64 {
	if irw.Data != nil {
		return irw.Data.GetTimeTaken()
	}
	return 0
}

// GetRequestId returns the request ID.
func (irw *InvokeResponseWrapper) GetRequestId() uint64 {
	if irw.Data != nil {
		return irw.Data.GetRequestId()
	}
	return 0
}

// ToJSON converts the wrapped data to JSON.
func (irw *InvokeResponseWrapper) ToJSON() (string, error) {
	if irw.Data == nil {
		return "", errors.New("no data to convert to JSON")
	}
	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
		Indent:          "  ",
		UseProtoNames:   true,
	}
	jsonData, err := marshaler.Marshal(irw.Data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// ToDict converts the wrapped data to a map.
func (irw *InvokeResponseWrapper) ToDict() (map[string]interface{}, error) {
	jsonString, err := irw.ToJSON()
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetData returns a list of Content objects.
func (irw *InvokeResponseWrapper) GetData() ([]*Content, error) {
	if irw.Data == nil {
		return nil, errors.New("no data available")
	}

	var contentList []*Content
	for _, cnt := range irw.Data.Responses {
		contentList = append(contentList, NewContent(cnt))
	}
	return contentList, nil
}

// GetMetadata returns the metadata as a map.
func (irw *InvokeResponseWrapper) GetMetadata() (map[string]interface{}, error) {
	if irw.Data == nil || irw.Data.Meta == nil {
		return nil, errors.New("no metadata available")
	}

	metaJSON, err := protojson.Marshal(irw.Data.Meta)
	if err != nil {
		return nil, err
	}

	var metadata map[string]interface{}
	if err := json.Unmarshal(metaJSON, &metadata); err != nil {
		return nil, err
	}

	return metadata, nil
}

// GetMetrics returns a list of Metric objects.
func (irw *InvokeResponseWrapper) GetMetrics() ([]*Metric, error) {
	if irw.Data == nil {
		return nil, errors.New("no data available")
	}

	var metricList []*Metric
	for _, mtr := range irw.Data.Metrics {
		metricList = append(metricList, NewMetric(mtr))
	}
	return metricList, nil
}

// GetCode returns the response code.
func (irw *InvokeResponseWrapper) GetCode() int32 {
	return irw.Code
}

// IsSuccess returns true if the operation was successful.
func (irw *InvokeResponseWrapper) IsSuccess() bool {
	return irw.Success
}

// IsError returns true if there was an error.
func (irw *InvokeResponseWrapper) IsError() bool {
	return !irw.IsSuccess()
}

// GetErrorCode returns the error code.
func (irw *InvokeResponseWrapper) GetErrorCode() int64 {
	if irw.Error != nil {
		return irw.GetErrorCode()
	}
	return 0
}

// GetError returns the error object.
func (irw *InvokeResponseWrapper) GetError() *lexatic_backend.InvokerError {
	return irw.Error
}

// GetErrorMessage returns the error message.
func (irw *InvokeResponseWrapper) GetErrorMessage() string {
	if irw.Error != nil {
		return irw.Error.ErrorMessage
	}
	log.Println("No error message found in response")
	return "No error message found in response"
}

// GetHumanErrorMessage returns the human-readable error message.
func (irw *InvokeResponseWrapper) GetHumanErrorMessage() string {
	if irw.Error != nil {
		return irw.Error.HumanMessage
	}
	log.Println("No human error message found in response")
	return "No human error message found in response"
}
