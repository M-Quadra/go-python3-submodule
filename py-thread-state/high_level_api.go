package pythreadstate

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import (
	python "github.com/M-Quadra/go-python3-submodule/v9"
)

// Get PyThreadState_Get
func Get() *python.PyThreadState {
	return toThreadState(C.PyThreadState_Get())
}

// Swap PyThreadState_Swap
func Swap(tstate *python.PyThreadState) *python.PyThreadState {
	return toThreadState(C.PyThreadState_Swap(toCThreadState(tstate)))
}
