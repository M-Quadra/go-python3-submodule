package pybytearray

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule/v11"
	"github.com/M-Quadra/go-python3-submodule/v11/py"
)

// FromObject PyByteArray_FromObject
func FromObject(o *python.PyObject) *python.PyObject {
	return toObject(C.PyByteArray_FromObject(toC(o)))
}

// FromString PyByteArray_FromStringAndSize
func FromString(str string) *python.PyObject {
	strC := C.CString(str)
	defer C.free(unsafe.Pointer(strC))

	return toObject(C.PyByteArray_FromStringAndSize(strC, (C.Py_ssize_t)(len(str))))
}

// Concat PyByteArray_Concat
func Concat(a, b *python.PyObject) *python.PyObject {
	if a == nil && b == nil {
		return nil
	}

	if a == nil {
		a = FromString("")
		defer py.DecRef(a)
	} else if b == nil {
		b = FromString("")
		defer py.DecRef(b)
	}
	return toObject(C.PyByteArray_Concat(toC(a), toC(b)))
}

// Size PyByteArray_Size
func Size(bytearray *python.PyObject) int {
	return int(C.PyByteArray_Size(toC(bytearray)))
}

// AsString PyByteArray_AsString
func AsString(bytearray *python.PyObject) string {
	return C.GoStringN(C.PyByteArray_AsString(toC(bytearray)), C.int(C.PyByteArray_Size(toC(bytearray))))
}

// Resize PyByteArray_Resize
func Resize(bytearray *python.PyObject, len int) int {
	if bytearray == nil {
		return 0
	}

	return int(C.PyByteArray_Resize(toC(bytearray), C.Py_ssize_t(len)))
}
