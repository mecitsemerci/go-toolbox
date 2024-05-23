package jsonconv

import (
	"encoding/json"
	"errors"
)

var (
	// EmptyStr is an empty string constant.
	EmptyStr string
	// Marshal is an alias for json.Marshal function.
	Marshal = json.Marshal
	// Unmarshal is an alias for json.Unmarshal function.
	Unmarshal = json.Unmarshal
)

// Serialize converts a Go value to JSON.
//
// Parameters:
//
//	value: The Go value to be converted.
//
// Returns:
//
//	A string representing the JSON value, or an empty string and an error if conversion fails.
func Serialize[T any](value T) (string, error) {
	bytes, err := Marshal(value)
	if err != nil {
		return EmptyStr, err
	}
	return string(bytes), nil
}

// SerializeCool converts a Go value to JSON without error checking.
//
// Parameters:
//
//	value: The Go value to be converted.
//
// Returns:
//
//	A string representing the JSON value.
//
// Note: This function does not handle errors, so it should be used when error checking is not necessary.
func SerializeCool[T any](value T) string {
	result, _ := Serialize(value)
	return result
}

// Deserialize converts a JSON string to a Go value.
//
// Parameters:
//
//	data: The JSON string to be converted.
//
// Returns:
//
//	The Go value represented by the JSON string, or an error if conversion fails.
func Deserialize[T any](data string) (T, error) {
	var result T
	if err := Unmarshal([]byte(data), &result); err != nil {
		return result, err
	}
	return result, nil
}

// DeserializeInto converts a JSON string to a Go value into a provided variable.
//
// Parameters:
//
//	data: The JSON string to be converted.
//	result: A pointer to the variable where the converted Go value will be stored.
//
// Returns:
//
//	An error if the conversion fails or if the result pointer is nil.
func DeserializeInto[T any](data string, result *T) error {
	if result == nil {
		return errors.New("result cannot be nil")
	}
	return json.Unmarshal([]byte(data), result)
}
