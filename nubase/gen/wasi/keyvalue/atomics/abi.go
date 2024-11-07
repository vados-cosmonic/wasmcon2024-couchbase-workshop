// Code generated by wit-bindgen-go. DO NOT EDIT.

package atomics

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/vados-cosmonic/wasmcon2024-couchbase-workshop/gen/wasi/keyvalue/store"
	"unsafe"
)

// ErrorShape is used for storage in variant or result types.
type ErrorShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(store.Error{})]byte
}
