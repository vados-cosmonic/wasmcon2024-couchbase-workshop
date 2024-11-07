// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package logging represents the imported interface "wasi:logging/logging@0.1.0-draft".
package logging

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// Level represents the enum "wasi:logging/logging@0.1.0-draft#level".
//
//	enum level {
//		trace,
//		debug,
//		info,
//		warn,
//		error,
//		critical
//	}
type Level uint8

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelCritical
)

var stringsLevel = [6]string{
	"trace",
	"debug",
	"info",
	"warn",
	"error",
	"critical",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e Level) String() string {
	return stringsLevel[e]
}

// Log represents the imported function "log".
//
//	log: func(level: level, context: string, message: string)
//
//go:nosplit
func Log(level Level, context string, message string) {
	level0 := (uint32)(level)
	context0, context1 := cm.LowerString(context)
	message0, message1 := cm.LowerString(message)
	wasmimport_Log((uint32)(level0), (*uint8)(context0), (uint32)(context1), (*uint8)(message0), (uint32)(message1))
	return
}
