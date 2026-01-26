// Copyright (c) 2023-2025 RapidaAI
// Author: Prashant Srivastav <prashant@rapida.ai>

package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestOrganizationKnowledgeCollection(t *testing.T) {
	tests := []struct {
		name        string
		orgId       uint64
		projectId   uint64
		knowledgeId uint64
		expected    string
	}{
		{
			name:        "basic case",
			orgId:       1,
			projectId:   2,
			knowledgeId: 3,
			expected:    "1__2__3",
		},
		{
			name:        "zeros",
			orgId:       0,
			projectId:   0,
			knowledgeId: 0,
			expected:    "0__0__0",
		},
		{
			name:        "large numbers",
			orgId:       18446744073709551615,
			projectId:   18446744073709551614,
			knowledgeId: 18446744073709551613,
			expected:    "18446744073709551615__18446744073709551614__18446744073709551613",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := OrganizationKnowledgeCollection(tt.orgId, tt.projectId, tt.knowledgeId)
			if result != tt.expected {
				t.Errorf("OrganizationKnowledgeCollection() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestOrganizationObjectPrefix(t *testing.T) {
	tests := []struct {
		name      string
		orgId     uint64
		projectId uint64
		prefix    string
		expected  string
	}{
		{
			name:      "basic case",
			orgId:     1,
			projectId: 2,
			prefix:    "documents",
			expected:  "1/2/documents",
		},
		{
			name:      "empty prefix",
			orgId:     1,
			projectId: 2,
			prefix:    "",
			expected:  "1/2/",
		},
		{
			name:      "nested prefix",
			orgId:     1,
			projectId: 2,
			prefix:    "path/to/object",
			expected:  "1/2/path/to/object",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := OrganizationObjectPrefix(tt.orgId, tt.projectId, tt.prefix)
			if result != tt.expected {
				t.Errorf("OrganizationObjectPrefix() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestPtr(t *testing.T) {
	t.Run("string pointer", func(t *testing.T) {
		value := "test"
		result := Ptr(value)
		if *result != value {
			t.Errorf("Ptr() = %q, want %q", *result, value)
		}
	})

	t.Run("int pointer", func(t *testing.T) {
		value := 42
		result := Ptr(value)
		if *result != value {
			t.Errorf("Ptr() = %d, want %d", *result, value)
		}
	})

	t.Run("bool pointer", func(t *testing.T) {
		value := true
		result := Ptr(value)
		if *result != value {
			t.Errorf("Ptr() = %v, want %v", *result, value)
		}
	})

	t.Run("struct pointer", func(t *testing.T) {
		type testStruct struct {
			Name string
			Age  int
		}
		value := testStruct{Name: "test", Age: 25}
		result := Ptr(value)
		if *result != value {
			t.Errorf("Ptr() = %v, want %v", *result, value)
		}
	})
}

func TestUnPtr(t *testing.T) {
	t.Run("non-nil string pointer", func(t *testing.T) {
		value := "test"
		ptr := &value
		result := UnPtr(ptr)
		if result != value {
			t.Errorf("UnPtr() = %q, want %q", result, value)
		}
	})

	t.Run("nil string pointer", func(t *testing.T) {
		var ptr *string
		result := UnPtr(ptr)
		if result != "" {
			t.Errorf("UnPtr() = %q, want empty string", result)
		}
	})

	t.Run("nil int pointer", func(t *testing.T) {
		var ptr *int
		result := UnPtr(ptr)
		if result != 0 {
			t.Errorf("UnPtr() = %d, want 0", result)
		}
	})

	t.Run("nil bool pointer", func(t *testing.T) {
		var ptr *bool
		result := UnPtr(ptr)
		if result != false {
			t.Errorf("UnPtr() = %v, want false", result)
		}
	})
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		name     string
		value    uint64
		expected string
	}{
		{
			name:     "zero",
			value:    0,
			expected: "0",
		},
		{
			name:     "positive number",
			value:    12345,
			expected: "12345",
		},
		{
			name:     "max uint64",
			value:    18446744073709551615,
			expected: "18446744073709551615",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntToString(tt.value)
			if result != tt.expected {
				t.Errorf("IntToString() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestDurationToString(t *testing.T) {
	tests := []struct {
		name     string
		value    time.Duration
		expected string
	}{
		{
			name:     "zero duration",
			value:    0,
			expected: "0",
		},
		{
			name:     "one second",
			value:    time.Second,
			expected: "1000000000",
		},
		{
			name:     "one minute",
			value:    time.Minute,
			expected: "60000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DurationToString(tt.value)
			if result != tt.expected {
				t.Errorf("DurationToString() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestToJson(t *testing.T) {
	t.Run("struct to json map", func(t *testing.T) {
		type testStruct struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		obj := testStruct{Name: "test", Age: 25}
		result := ToJson(obj)
		if result["name"] != "test" {
			t.Errorf("ToJson()['name'] = %v, want 'test'", result["name"])
		}
		if result["age"] != float64(25) {
			t.Errorf("ToJson()['age'] = %v, want 25", result["age"])
		}
	})

	t.Run("nil input", func(t *testing.T) {
		result := ToJson(nil)
		// ToJson returns an empty map (not nil) when json marshal/unmarshal succeeds
		// but for nil input, the result will be nil as json.Marshal(nil) produces "null"
		// which unmarshal into map returns nil
		if result != nil && len(result) != 0 {
			t.Errorf("ToJson(nil) should return nil or empty map, got %v", result)
		}
	})

	t.Run("map input", func(t *testing.T) {
		obj := map[string]interface{}{"key": "value"}
		result := ToJson(obj)
		if result["key"] != "value" {
			t.Errorf("ToJson()['key'] = %v, want 'value'", result["key"])
		}
	})
}

func TestSerialize(t *testing.T) {
	t.Run("basic map", func(t *testing.T) {
		data := map[string]interface{}{
			"key1": "value1",
			"key2": 123,
		}
		result, err := Serialize(data)
		if err != nil {
			t.Errorf("Serialize() error = %v", err)
		}
		if result == nil {
			t.Error("Serialize() returned nil")
		}
	})

	t.Run("map with error value", func(t *testing.T) {
		data := map[string]interface{}{
			"key1":  "value1",
			"error": error(nil),
		}
		result, err := Serialize(data)
		if err != nil {
			t.Errorf("Serialize() error = %v", err)
		}
		if result == nil {
			t.Error("Serialize() returned nil")
		}
	})

	t.Run("empty map", func(t *testing.T) {
		data := map[string]interface{}{}
		result, err := Serialize(data)
		if err != nil {
			t.Errorf("Serialize() error = %v", err)
		}
		expected := []byte("{}")
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Serialize() = %s, want %s", result, expected)
		}
	})
}
