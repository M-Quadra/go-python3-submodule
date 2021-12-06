package pybytearray

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import python "github.com/M-Quadra/go-python3-submodule/v9"

// Check PyByteArray_Check
func Check(o *python.PyObject) bool {
	if o == nil {
		return false
	}

	return C.cgo_PyByteArray_Check(toC(o)) != 0
}

// CheckExact PyByteArray_CheckExact
func CheckExact(o *python.PyObject) bool {
	if o == nil {
		return false
	}

	return C.cgo_PyByteArray_CheckExact(toC(o)) != 0
}
