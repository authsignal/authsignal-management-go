package authsignal

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"
)

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

func marshalClearingMembers(value any, clear []string, clearable ...string) ([]byte, error) {
	encoded, err := json.Marshal(value)
	if err != nil || len(clear) == 0 {
		return encoded, err
	}

	var setMembers map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &setMembers); err != nil {
		return nil, err
	}

	encoded = encoded[:len(encoded)-1]

	for index, name := range clear {
		if !slices.Contains(clearable, name) {
			return nil, fmt.Errorf("cannot clear %q, the members that can be cleared are %s", name, strings.Join(clearable, ", "))
		}

		if _, isSet := setMembers[name]; isSet {
			return nil, fmt.Errorf("cannot clear %q because it is also set to a value", name)
		}

		if slices.Contains(clear[:index], name) {
			continue
		}

		if len(encoded) > 1 {
			encoded = append(encoded, ',')
		}

		encoded = fmt.Appendf(encoded, "%q:null", name)
	}

	return append(encoded, '}'), nil
}
