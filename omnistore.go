// Package omnistore is designed to store anything and subsequently retrieve it.
//
//	omnistore.Set[float64]("pi", 3.14)
//	pi := omnistore.Get[float64]("pi")
//
//	omnistore.Set[int]("meaningoflife", 42)
//	mol := omnistore.Get[int]("meaningoflife")
package omnistore

import (
	"errors"
	"fmt"
	"sync"
)

var internalStore = make(map[string]any)

var (
	ErrValueNotFound = errors.New("value not found")
	ErrWrongType     = errors.New("wrong type")
)

// Set can store any value via key/value pair.
// This function is concurrency safe.
func Set[T any](key string, value T) {
	var lock sync.Mutex
	defer lock.Unlock()
	lock.Lock()

	internalStore[key] = value
}

// StringerSet can store any value via key/value pair as long as the key implements fmt.Stringer.
// This can be handy if you have a custom type enum you'd like to use rather than string values.
// This function is concurrency safe.
func StringerSet[K fmt.Stringer, V any](key K, value V) {
	Set[V](key.String(), value)
}

// Get will return any value via key name if the internal store has the value.
// If no value is found, the zero value of the value type will be returned.
// This function is concurrency safe.
func Get[T any](key string) T {
	var lock sync.Mutex
	defer lock.Unlock()
	lock.Lock()

	var zeroValue T
	if v, ok := internalStore[key].(T); ok {
		return v
	}

	return zeroValue
}

// GetE will return any value via key name if the internal store has the value or an error.
// This function is concurrency safe.
func GetE[T any](key string) (T, error) {
	var lock sync.Mutex
	defer lock.Unlock()
	lock.Lock()

	var zeroValue T
	if value, ok := internalStore[key]; ok {
		if v, ok := value.(T); ok {
			return v, nil
		}
		return zeroValue, fmt.Errorf("type %T: %w", zeroValue, ErrWrongType)
	}

	return zeroValue, fmt.Errorf("value of key %s: %w", key, ErrValueNotFound)
}

// StringerGet will return any value of a key that implements fmt.Stringer if the internal store has the value.
// If no value is found, the zero value of the value type will be returned.
// This can be handy if you have a custom type enum you'd like to use rather than string values.
// This function is concurrency safe.
func StringerGet[K fmt.Stringer, V any](key K) V {
	return Get[V](key.String())
}

// StringerGetE will return any value of a key that implements fmt.Stringer if the internal store has the value.
// If no value is found an error is returned.
// This can be handy if you have a custom type enum you'd like to use rather than string values.
// This function is concurrency safe.
func StringerGetE[K fmt.Stringer, V any](key K) (V, error) {
	return GetE[V](key.String())
}
