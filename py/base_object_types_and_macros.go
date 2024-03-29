package py

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import python "github.com/M-Quadra/go-python3-submodule/v11"

// Py_Is

// IsNone Py_IsNone
func IsNone(o *python.PyObject) bool {
	return C.Py_IsNone(toC(o)) != 0
}

// Py_TYPE

// Py_IS_TYPE

// Py_SET_TYPE

// RefCnt Py_REFCNT
func RefCnt(o *python.PyObject) int {
	return int(C.cgo_Py_REFCNT(toC(o)))
}

// Py_SET_REFCNT

// Py_SIZE

// Py_SET_SIZE

// PyObject_HEAD_INIT

// PyVarObject_HEAD_INIT
