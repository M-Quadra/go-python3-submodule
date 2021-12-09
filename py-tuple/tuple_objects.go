package pytuple

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import (
	python "github.com/M-Quadra/go-python3-submodule/v8"
)

// Check PyTuple_Check
func Check(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyTuple_Check(toC(p)) != 0
}

// CheckExact PyTuple_CheckExact
func CheckExact(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyTuple_CheckExact(toC(p)) != 0
}

// New PyTuple_New
func New(len int) *python.PyObject {
	return toObject(C.PyTuple_New(C.Py_ssize_t(len)))
}

// PyTuple_Pack

// Size PyTuple_Size
func Size(p *python.PyObject) int {
	if p == nil {
		return 0
	}

	return int(C.PyTuple_Size(toC(p)))
}

// PyTuple_GET_SIZE

// GetItem PyTuple_GetItem
func GetItem(p *python.PyObject, pos int) *python.PyObject {
	if p == nil {
		return nil
	}

	return toObject(C.PyTuple_GetItem(toC(p), (C.Py_ssize_t)(pos)))
}

// PyTuple_GET_ITEM

// GetSlice PyTuple_GetSlice
func GetSlice(p *python.PyObject, low, high int) *python.PyObject {
	return toObject(C.PyTuple_GetSlice(toC(p), C.Py_ssize_t(low), C.Py_ssize_t(high)))
}

// SetItem PyTuple_SetItem
func SetItem(p *python.PyObject, pos int, o *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.PyTuple_SetItem(toC(p), C.Py_ssize_t(pos), toC(o)) == 0
}

// PyTuple_SET_ITEM

// _PyTuple_Resize
