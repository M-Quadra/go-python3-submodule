package pylist

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import python "github.com/M-Quadra/go-python3-submodule/v9"

// Check PyList_Check
func Check(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyList_Check(toC(p)) != 0
}

// CheckExact PyList_CheckExact
func CheckExact(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyList_CheckExact(toC(p)) != 0
}

// New PyList_New
func New(len int) *python.PyObject {
	return toObject(C.PyList_New(C.Py_ssize_t(len)))
}

// Size PyList_Size
func Size(list *python.PyObject) int {
	if list == nil {
		return 0
	}

	return int(C.PyList_Size(toC(list)))
}

// PyList_GET_SIZE

// GetItem PyList_GetItem
func GetItem(list *python.PyObject, index int) *python.PyObject {
	if list == nil {
		return nil
	}

	return toObject(C.PyList_GetItem(toC(list), C.Py_ssize_t(index)))
}

// PyList_GET_ITEM

// SetItem PyList_SetItem
func SetItem(list *python.PyObject, index int, item *python.PyObject) bool {
	if list == nil {
		return false
	}

	return C.PyList_SetItem(toC(list), C.Py_ssize_t(index), toC(item)) == 0
}

// PyList_SET_ITEM

// Insert PyList_Insert
func Insert(list *python.PyObject, index int, item *python.PyObject) bool {
	if list == nil {
		return false
	}

	return C.PyList_Insert(toC(list), C.Py_ssize_t(index), toC(item)) == 0
}

// Append PyList_Append
func Append(list *python.PyObject, item *python.PyObject) bool {
	if list == nil {
		return false
	}

	return C.PyList_Append(toC(list), toC(item)) == 0
}

// GetSlice PyList_GetSlice
func GetSlice(list *python.PyObject, low, high int) *python.PyObject {
	if list == nil {
		return nil
	}

	return toObject(C.PyList_GetSlice(toC(list), C.Py_ssize_t(low), C.Py_ssize_t(high)))
}

// SetSlice PyList_SetSlice
func SetSlice(list *python.PyObject, low, high int, itemlist *python.PyObject) bool {
	return C.PyList_SetSlice(toC(list), C.Py_ssize_t(low), C.Py_ssize_t(high), toC(itemlist)) == 0
}

// Sort PyList_Sort
func Sort(list *python.PyObject) bool {
	return C.PyList_Sort(toC(list)) == 0
}

// Reverse PyList_Reverse
func Reverse(list *python.PyObject) bool {
	return C.PyList_Reverse(toC(list)) == 0
}

// AsTuple PyList_AsTuple
func AsTuple(list *python.PyObject) *python.PyObject {
	return toObject(C.PyList_AsTuple(toC(list)))
}
