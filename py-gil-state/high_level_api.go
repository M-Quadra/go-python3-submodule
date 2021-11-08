package pygilstate

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule"
)

// Ensure PyGILState_Ensure
func Ensure() python.PyGILState {
	return toGILState(C.PyGILState_Ensure())
}

// Release PyGILState_Release
func Release(state python.PyGILState) {
	C.PyGILState_Release(toCGILState(state))
}

// GetThisThreadState PyGILState_GetThisThreadState
func GetThisThreadState() *python.PyThreadState {
	return (*python.PyThreadState)(unsafe.Pointer(C.PyGILState_GetThisThreadState()))
}

// Check PyGILState_Check
func Check() bool {
	return C.PyGILState_Check() == 1
}
