package validation

import (
	"errors"
	"reflect"

	"golang.org/x/exp/constraints"
)

// isEmpty checks if a given object is empty.
// It supports various types including arrays, channels, maps, slices, and pointers.
// For pointers, it recursively checks the referenced object.
// For other types, it compares the object with its zero value.
func isEmpty(obj any) bool {
	if obj == nil {
		return true
	}

	objValue := reflect.ValueOf(obj)

	switch objValue.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		ref := objValue.Elem().Interface()
		return isEmpty(ref)
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(obj, zero.Interface())
	}
}

func requireThat(checkFn func() bool) error {
	if !checkFn() {
		return errors.New("condition not met")
	}
	return nil
}

// NotEmpty is a generic function that returns a validation function for checking if a field is not empty.
// It takes a function getField that retrieves the field value from an entity.
// The returned validation function returns an error if the field is empty.
func NotEmpty[T any, N any](getField func(T) N) func(T) error {
	return func(entity T) error {
		if isEmpty(getField(entity)) {
			return errors.New("field cannot be empty")
		}
		return nil
	}
}

// MaxLength is a generic function that returns a validation function for checking if a field length does not exceed a maximum length.
// It takes a function getField that retrieves the field value from an entity and an integer maxLength.
// The returned validation function returns an error if the field length exceeds maxLength.
func MaxLength[T any](getField func(T) string, maxLength int) func(T) error {
	return func(entity T) error {
		if len(getField(entity)) > maxLength {
			return errors.New("field exceeds maximum length")
		}
		return nil
	}
}

// MinLength is a generic function that returns a validation function for checking if a field length is not below a minimum length.
// It takes a function getField that retrieves the field value from an entity and an integer minLength.
// The returned validation function returns an error if the field length is below minLength.
func MinLength[T any](getField func(T) string, minLength int) func(T) error {
	return func(entity T) error {
		if len(getField(entity)) < minLength {
			return errors.New("field is below minimum length")
		}
		return nil
	}
}

// IsInRange is a generic function that returns a validation function for checking if a field value is within a specified range.
// It takes a function getField that retrieves the field value from an entity, and two values min and max of type N, which must be ordered.
// The returned validation function returns an error if the field value is outside the range [min, max].
func IsInRange[T any, N constraints.Ordered](getField func(T) N, min, max N) func(T) error {
	return func(entity T) error {
		val := getField(entity)
		if val < min || val > max {
			return errors.New("field is out of range")
		}
		return nil
	}
}
