// Code generated by wit-bindgen-go. DO NOT EDIT.

package logging

// This file contains wasmimport and wasmexport declarations for "wasi:logging@0.1.0-draft".

//go:wasmimport wasi:logging/logging@0.1.0-draft log
//go:noescape
func wasmimport_Log(level0 uint32, context0 *uint8, context1 uint32, message0 *uint8, message1 uint32)