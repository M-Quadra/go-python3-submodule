package pyerr

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import python "github.com/M-Quadra/go-python3-submodule/v10"

// Clear PyErr_Clear
func Clear() {
	C.PyErr_Clear()
}

// PrintEx PyErr_PrintEx
func PrintEx(setSysLastVars bool) {
	if setSysLastVars {
		C.PyErr_PrintEx(1)
	} else {
		C.PyErr_PrintEx(0)
	}
}

// Print PyErr_Print
func Print() {
	C.PyErr_Print()
}

// WriteUnraisable PyErr_WriteUnraisable
func WriteUnraisable(obj *python.PyObject) {
	C.PyErr_WriteUnraisable(toC(obj))
}
