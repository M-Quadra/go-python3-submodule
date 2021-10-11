package pymodule

import (
	/*
		#cgo pkg-config: python3-embed
		#include "Python.h"
	*/
	"C"
)

import (
	"python"
	"unsafe"
)

func toC(obj *python.PyObject) *C.PyObject {
	return (*C.PyObject)(unsafe.Pointer(obj))
}

func toObject(obj *C.PyObject) *python.PyObject {
	return (*python.PyObject)(unsafe.Pointer(obj))
}
