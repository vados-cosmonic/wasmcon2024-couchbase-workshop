// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package atomics represents the imported interface "wrpc:keyvalue/atomics@0.2.0-draft".
//
// A keyvalue interface that provides atomic operations.
//
// Atomic operations are single, indivisible operations. When a fault causes an atomic
// operation to
// fail, it will appear to the invoker of the atomic operation that the action either
// completed
// successfully or did nothing at all.
//
// Please note that this interface is bare functions that take a reference to a bucket.
// This is to
// get around the current lack of a way to "extend" a resource with additional methods
// inside of
// wit. Future version of the interface will instead extend these methods on the base
// `bucket`
// resource.
package atomics

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/vados-cosmonic/wasmcon2024-couchbase-workshop/gen/wrpc/keyvalue/store"
)

// Error represents the type alias "wrpc:keyvalue/atomics@0.2.0-draft#error".
//
// See [store.Error] for more information.
type Error = store.Error

// Increment represents the imported function "increment".
//
// Atomically increment the value associated with the key in the store by the given
// delta. It
// returns the new value.
//
// If the key does not exist in the store, it creates a new key-value pair with the
// value set
// to the given delta.
//
// If any other error occurs, it returns an `Err(error)`.
//
//	increment: func(bucket: string, key: string, delta: u64) -> result<u64, error>
//
//go:nosplit
func Increment(bucket string, key string, delta uint64) (result cm.Result[ErrorShape, uint64, Error]) {
	bucket0, bucket1 := cm.LowerString(bucket)
	key0, key1 := cm.LowerString(key)
	delta0 := (uint64)(delta)
	wasmimport_Increment((*uint8)(bucket0), (uint32)(bucket1), (*uint8)(key0), (uint32)(key1), (uint64)(delta0), &result)
	return
}
