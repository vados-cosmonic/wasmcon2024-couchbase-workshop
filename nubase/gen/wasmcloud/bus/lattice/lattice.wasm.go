// Code generated by wit-bindgen-go. DO NOT EDIT.

package lattice

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasmcloud:bus@1.0.0".

//go:wasmimport wasmcloud:bus/lattice@1.0.0 [resource-drop]call-target-interface
//go:noescape
func wasmimport_CallTargetInterfaceResourceDrop(self0 uint32)

//go:wasmimport wasmcloud:bus/lattice@1.0.0 [constructor]call-target-interface
//go:noescape
func wasmimport_NewCallTargetInterface(namespace0 *uint8, namespace1 uint32, package0 *uint8, package1 uint32, interface0 *uint8, interface1 uint32) (result0 uint32)

//go:wasmimport wasmcloud:bus/lattice@1.0.0 set-link-name
//go:noescape
func wasmimport_SetLinkName(name0 *uint8, name1 uint32, interfaces0 *CallTargetInterface, interfaces1 uint32, result *cm.Result[string, struct{}, string])
