package pymodule

import (
	/*
		#cgo pkg-config: python3-embed
		#include "Python.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule"
)

func toC(obj *python.PyObject) *C.PyObject {
	return (*C.PyObject)(unsafe.Pointer(obj))
}

func toObject(obj *C.PyObject) *python.PyObject {
	return (*python.PyObject)(unsafe.Pointer(obj))
}
