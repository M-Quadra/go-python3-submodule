package pytype

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule"
)

// Type PyType_Type
var Type = toObject((*C.PyObject)(unsafe.Pointer(&C.PyType_Type)))

// Check PyType_Check
func Check(o *python.PyObject) bool {
	if o == nil {
		return false
	}

	return C.cgo_PyType_Check(toC(o)) != 0
}

// CheckExact PyType_CheckExact
func CheckExact(o *python.PyObject) bool {
	if o == nil {
		return false
	}

	return C.cgo_PyType_CheckExact(toC(o)) != 0
}

// ClearCache PyType_ClearCache
func ClearCache() uint {
	return uint(C.PyType_ClearCache())
}

// PyType_GetFlags

// PyType_Modified

// PyType_HasFeature

// PyType_IS_GC

// PyType_IsSubtype

// PyType_GenericAlloc

// PyType_GenericNew

// PyType_Ready

// PyType_GetSlot

// PyType_GetModule

// PyType_GetModuleState
