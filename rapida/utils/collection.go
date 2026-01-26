// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"strings"
)

func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			if nestedMap, ok := v.(map[string]interface{}); ok {
				if existingNestedMap, exists := result[k].(map[string]interface{}); exists {
					result[k] = MergeMaps(existingNestedMap, nestedMap)
				} else {
					result[k] = MergeMaps(nestedMap)
				}
			} else {
				result[k] = v
			}
		}
	}
	return result
}

func GetCaseInsensitiveKeyValue(cfg map[string]string, key string) (string, bool) {
	if value, ok := cfg[key]; ok {
		return value, true
	}
	if value, ok := cfg[strings.ToUpper(key)]; ok {
		return value, true
	}
	return "", false
}

func EmbeddingToFloat64[T float32 | float64](embedding []T) []float64 {
	float64Embedding := make([]float64, len(embedding))
	for i, val := range embedding {
		float64Embedding[i] = float64(val)
	}
	return float64Embedding
}

func EmbeddingToFloat32[T float32 | float64](embedding []T) []float32 {
	float32Embedding := make([]float32, len(embedding))
	for i, val := range embedding {
		float32Embedding[i] = float32(val)
	}
	return float32Embedding
}

// Convert a slice of float32 to a byte array
func Float64SliceToByteArray(data []float64) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func EmbeddingToBase64(embedding []float64) string {
	byteArray, err := Float64SliceToByteArray(embedding)
	if err != nil {
		return ""
	}
	base64Str := base64.StdEncoding.EncodeToString(byteArray)
	return base64Str
}
