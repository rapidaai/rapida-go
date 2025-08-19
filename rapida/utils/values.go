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
 *  Author: Prashant Srivastav
 *
 */
package rapida_utils

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ConvertStringToAny converts a string to *anypb.Any
func StringToAny(value string) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	err := anypb.MarshalFrom(anyValue, wrapperspb.String(value), proto.MarshalOptions{})
	return anyValue, err
}

// ConvertAnyToString converts *anypb.Any to a string
func AnyToString(anyValue *anypb.Any) (string, error) {
	stringWrapper := &wrapperspb.StringValue{}
	err := anyValue.UnmarshalTo(stringWrapper)
	if err != nil {
		return "", err
	}
	return stringWrapper.GetValue(), nil
}

// ConvertFloat64ToAny converts a float64 to *anypb.Any
func Float64ToAny(value float64) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	err := anypb.MarshalFrom(anyValue, wrapperspb.Double(value), proto.MarshalOptions{})
	return anyValue, err
}

// ConvertAnyToFloat64 converts *anypb.Any to a float64
func AnyToFloat64(anyValue *anypb.Any) (float64, error) {
	floatWrapper := &wrapperspb.DoubleValue{}
	err := anyValue.UnmarshalTo(floatWrapper)
	if err != nil {
		return 0, err
	}
	return floatWrapper.GetValue(), nil
}

// ConvertFloat32ToAny converts a float32 to *anypb.Any
func Float32ToAny(value float32) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	err := anypb.MarshalFrom(anyValue, wrapperspb.Float(value), proto.MarshalOptions{})
	return anyValue, err
}

// ConvertAnyToFloat32 converts *anypb.Any to a float32
func AnyToFloat32(anyValue *anypb.Any) (float32, error) {
	floatWrapper := &wrapperspb.FloatValue{}
	err := anyValue.UnmarshalTo(floatWrapper)
	if err != nil {
		return 0, err
	}
	return floatWrapper.GetValue(), nil
}

// ConvertInt32ToAny converts an int32 to *anypb.Any
func Int32ToAny(value int32) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	err := anypb.MarshalFrom(anyValue, wrapperspb.Int32(value), proto.MarshalOptions{})
	return anyValue, err
}

// ConvertAnyToInt32 converts *anypb.Any to an int32
func AnyToInt32(anyValue *anypb.Any) (int32, error) {
	intWrapper := &wrapperspb.Int32Value{}
	err := anyValue.UnmarshalTo(intWrapper)
	if err != nil {
		return 0, err
	}
	return intWrapper.GetValue(), nil
}

// ConvertInt64ToAny converts an int64 to *anypb.Any
func Int64ToAny(value int64) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	err := anypb.MarshalFrom(anyValue, wrapperspb.Int64(value), proto.MarshalOptions{})
	return anyValue, err
}

// ConvertAnyToInt64 converts *anypb.Any to an int64
func AnyToInt64(anyValue *anypb.Any) (int64, error) {
	intWrapper := &wrapperspb.Int64Value{}
	err := anyValue.UnmarshalTo(intWrapper)
	if err != nil {
		return 0, err
	}
	return intWrapper.GetValue(), nil
}

// ConvertUInt32ToAny converts a uint32 to *anypb.Any
func UInt32ToAny(value uint32) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	err := anypb.MarshalFrom(anyValue, wrapperspb.UInt32(value), proto.MarshalOptions{})
	return anyValue, err
}

// ConvertAnyToUInt32 converts *anypb.Any to a uint32
func AnyToUInt32(anyValue *anypb.Any) (uint32, error) {
	uintWrapper := &wrapperspb.UInt32Value{}
	err := anyValue.UnmarshalTo(uintWrapper)
	if err != nil {
		return 0, err
	}
	return uintWrapper.GetValue(), nil
}

// ConvertUInt64ToAny converts a uint64 to *anypb.Any
func UInt64ToAny(value uint64) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	err := anypb.MarshalFrom(anyValue, wrapperspb.UInt64(value), proto.MarshalOptions{})
	return anyValue, err
}

// ConvertAnyToUInt64 converts *anypb.Any to a uint64
func AnyToUInt64(anyValue *anypb.Any) (uint64, error) {
	uintWrapper := &wrapperspb.UInt64Value{}
	err := anyValue.UnmarshalTo(uintWrapper)
	if err != nil {
		return 0, err
	}
	return uintWrapper.GetValue(), nil
}

// ConvertBoolToAny converts a bool to *anypb.Any
func BoolToAny(value bool) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	err := anypb.MarshalFrom(anyValue, wrapperspb.Bool(value), proto.MarshalOptions{})
	return anyValue, err
}

// ConvertAnyToBool converts *anypb.Any to a bool
func AnyToBool(anyValue *anypb.Any) (bool, error) {
	boolWrapper := &wrapperspb.BoolValue{}
	err := anyValue.UnmarshalTo(boolWrapper)
	if err != nil {
		return false, err
	}
	return boolWrapper.GetValue(), nil
}

// ConvertBytesToAny converts a []byte to *anypb.Any
func BytesToAny(value []byte) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	err := anypb.MarshalFrom(anyValue, wrapperspb.Bytes(value), proto.MarshalOptions{})
	return anyValue, err
}

// ConvertAnyToBytes converts *anypb.Any to a []byte
func AnyToBytes(anyValue *anypb.Any) ([]byte, error) {
	bytesWrapper := &wrapperspb.BytesValue{}
	err := anyValue.UnmarshalTo(bytesWrapper)
	if err != nil {
		return nil, err
	}
	return bytesWrapper.GetValue(), nil
}

// ConvertJSONToAny converts a JSON map to *anypb.Any
func JSONToAny(value map[string]interface{}) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	jsonData, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	err = anypb.MarshalFrom(anyValue, &wrapperspb.BytesValue{Value: jsonData}, proto.MarshalOptions{})
	if err == nil {
		return anyValue, nil
	}
	anyValue.TypeUrl = "type.googleapis.com/google.protobuf.Struct"
	anyValue.Value = jsonData
	return anyValue, nil
}

// ConvertAnyToJSON converts *anypb.Any to a JSON map
func AnyToJSON(anyValue *anypb.Any) (map[string]interface{}, error) {
	var value map[string]interface{}
	bytesWrapper := &wrapperspb.BytesValue{}
	err := anyValue.UnmarshalTo(bytesWrapper)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytesWrapper.Value, &value)
	if err != nil {
		return nil, err
	}
	return value, nil
}
