package authsignal

import "encoding/json"

type NullableJsonInput[T any] map[string]T

func SetValue[T any](value T) map[string]T {
	return map[string]T{"isSet": value}
}

func SetNull[T any](value T) map[string]T {
	return map[string]T{"isNull": value}
}

func (inputMap NullableJsonInput[T]) MarshalJSON() ([]byte, error) {
	_, isExplicitlySetToNull := inputMap["isNull"]

	if isExplicitlySetToNull {
		return []byte("null"), nil
	}

	return json.Marshal(inputMap["isSet"])
}

type ListJsonInput[T any] struct {
	values []T
}

func SetList[T any](values []T) *ListJsonInput[T] {
	return &ListJsonInput[T]{values: values}
}

func (input ListJsonInput[T]) MarshalJSON() ([]byte, error) {
	if input.values == nil {
		return []byte("[]"), nil
	}

	return json.Marshal(input.values)
}
