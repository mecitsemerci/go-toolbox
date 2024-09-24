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

// MustSerialize converts a Go value to JSON without error checking.
//
// MustSerialize is a convenience function that converts a Go value to JSON without error checking.
// It is intended for situations where error checking is not necessary.
//
// Parameters:
//
//	value: The Go value to be converted. This can be of any type that can be marshalled to JSON.
//
// Returns:
//
//	A string representing the JSON value. If an error occurs during the conversion,
//	this function will panic with the error.
//
// Note: This function does not handle errors, so it should be used when error checking is not necessary.
func MustSerialize[T any](value T) string {
	result, err := Serialize(value)
	if err != nil {
		panic(err)
	}
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
