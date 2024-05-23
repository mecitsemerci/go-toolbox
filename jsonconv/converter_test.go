package jsonconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerialize_ShouldSerializeStructWhenValidInput(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	testCases := []struct {
		name     string
		input    TestStruct
		expected string
	}{
		{
			name:     "Valid input",
			input:    TestStruct{Name: "John Doe", Age: 30},
			expected: `{"name":"John Doe","age":30}`,
		},
		{
			name:     "Empty input",
			input:    TestStruct{},
			expected: `{}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Serialize(tc.input)
			require.NoError(t, err)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestSerialize_ShouldReturnErrorWhenInvalidInput(t *testing.T) {
	type TestStruct struct {
		Name     string      `json:"name,omitempty"`
		Age      int         `json:"age,omitempty"`
		Messages chan string `json:"messages,omitempty"`
	}
	messages := make(chan string)
	testCases := []struct {
		name     string
		input    TestStruct
		expected string
	}{
		{
			name:     "Valid input",
			input:    TestStruct{Name: "John Doe", Age: 30, Messages: messages},
			expected: "",
		}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Serialize(tc.input)
			require.Error(t, err)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestSerializeCool_ShouldSerializeStructWhenValidInput(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	testCases := []struct {
		name     string
		input    TestStruct
		expected string
	}{
		{
			name:     "Valid input",
			input:    TestStruct{Name: "John Doe", Age: 30},
			expected: `{"name":"John Doe","age":30}`,
		},
		{
			name:     "Empty input",
			input:    TestStruct{},
			expected: `{}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SerializeCool(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestDeserialize_ShouldDeserializeStructWhenValidInput(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	testCases := []struct {
		name     string
		input    string
		expected TestStruct
	}{
		{
			name:     "Valid input",
			input:    `{"name":"John Doe","age":30}`,
			expected: TestStruct{Name: "John Doe", Age: 30},
		},
		{
			name:     "Empty input",
			input:    `{}`,
			expected: TestStruct{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Deserialize[TestStruct](tc.input)
			require.NoError(t, err)

			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestDeserialize_ShouldDeserializeStructWhenInvalidInput(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	testCases := []struct {
		name        string
		input       string
		expectedErr bool
	}{
		{
			name:        "Invalid input",
			input:       `{`,
			expectedErr: true,
		},
		{
			name:        "Invalid input 2",
			input:       ``,
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var expectedResult TestStruct
			result, err := Deserialize[TestStruct](tc.input)
			require.Error(t, err)

			assert.Equal(t, expectedResult, result)
		})
	}
}

func TestDeserializeInto_ShouldDeserializeStructWhenValidInput(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name     string
		input    string
		expected TestStruct
	}{
		{
			name:     "Valid input",
			input:    `{"name":"John Doe","age":30}`,
			expected: TestStruct{Name: "John Doe", Age: 30},
		},
		{
			name:     "Empty input",
			input:    `{}`,
			expected: TestStruct{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result TestStruct
			err := DeserializeInto(tc.input, &result)
			require.NoError(t, err)

			assert.Equal(t, tc.expected, result)
		})
	}

	t.Run("Nil result pointer", func(t *testing.T) {
		var result *TestStruct
		err := DeserializeInto(`{}`, result)
		require.Error(t, err)
		expectedError := "result cannot be nil"
		assert.Equal(t, expectedError, err.Error())

	})
}
