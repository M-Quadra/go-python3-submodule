package pybool

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import python "github.com/M-Quadra/go-python3-submodule/v11"

// Check PyBool_Check
func Check(o *python.PyObject) bool {
	if o == nil {
		return false
	}

	return C.cgo_PyBool_Check(toC(o)) == 0
}

// FromInt PyBool_FromLong
func FromInt(v int) *python.PyObject {
	return toObject(C.PyBool_FromLong(C.long(v)))
}
