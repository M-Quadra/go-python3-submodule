package pythreadstate

import (
	/*
		#cgo pkg-config: python3-embed
		#include "Python.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule/v8"
)

func toCThreadState(tstate *python.PyThreadState) *C.PyThreadState {
	return (*C.PyThreadState)(unsafe.Pointer(tstate))
}

func toThreadState(tstate *C.PyThreadState) *python.PyThreadState {
	return (*python.PyThreadState)(unsafe.Pointer(tstate))
}
