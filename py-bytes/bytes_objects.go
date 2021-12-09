package pybytes

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule/v8"
)

// Check PyBytes_Check
func Check(o *python.PyObject) bool {
	if o == nil {
		return false
	}

	return C.cgo_PyBytes_Check(toC(o)) != 0
}

// CheckExact PyBytes_CheckExact
func CheckExact(o *python.PyObject) bool {
	objC := toC(o)
	if objC == nil {
		return false
	}

	return C.cgo_PyBytes_CheckExact(objC) != 0
}

// FromString PyBytes_FromString
func FromString(v string) *python.PyObject {
	vC := C.CString(v)
	defer C.free(unsafe.Pointer(vC))

	return toObject(C.PyBytes_FromString(vC))
}

// PyBytes_FromStringAndSize

// PyBytes_FromFormat

// PyBytes_FromFormatV

// FromObject PyBytes_FromObject
func FromObject(o *python.PyObject) *python.PyObject {
	return toObject(C.PyBytes_FromObject(toC(o)))
}

// Size PyBytes_Size
func Size(o *python.PyObject) int {
	return int(C.PyBytes_Size(toC(o)))
}

// PyBytes_GET_SIZE

// AsString PyBytes_AsString
func AsString(o *python.PyObject) string {
	oC := toC(o)
	return C.GoStringN(C.PyBytes_AsString(oC), (C.int)(C.PyBytes_Size(oC)))
}

// PyBytes_AS_STRING

// PyBytes_AsStringAndSize

// Concat PyBytes_Concat
func Concat(bytes, newpart *python.PyObject) *python.PyObject {
	bytesC := toC(bytes)
	C.PyBytes_Concat(&bytesC, toC(newpart))
	return toObject(bytesC)
}

// ConcatAndDel PyBytes_ConcatAndDel
func ConcatAndDel(bytes, newpart *python.PyObject) *python.PyObject {
	bytesC := toC(bytes)
	C.PyBytes_ConcatAndDel(&bytesC, toC(newpart))
	return toObject(bytesC)
}

// _PyBytes_Resize
