package py

import (
	/*
		#include "Python.h"
	*/
	"C"
)

// Initialize Py_Initialize
func Initialize() {
	C.Py_Initialize()
}

// InitializeEx Py_InitializeEx
func InitializeEx(initsigs bool) {
	if initsigs {
		C.Py_InitializeEx(1)
	} else {
		C.Py_InitializeEx(0)
	}
}

// IsInitialized Py_IsInitialized
func IsInitialized() bool {
	return C.Py_IsInitialized() != 0
}

// FinalizeEx Py_FinalizeEx
func FinalizeEx() int {
	return int(C.Py_FinalizeEx())
}

// Finalize Py_Finalize
func Finalize() {
	C.Py_Finalize()
}
