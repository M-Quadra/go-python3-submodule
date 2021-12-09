package py

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

// EnterRecursiveCall Py_EnterRecursiveCall
func EnterRecursiveCall(where string) int {
	whereC := C.CString(where)
	defer C.free(unsafe.Pointer(whereC))

	return int(C.cgo_Py_EnterRecursiveCall(whereC))
}

// LeaveRecursiveCall Py_LeaveRecursiveCall
func LeaveRecursiveCall() {
	C.cgo_Py_LeaveRecursiveCall()
}

// ReprEnter Py_ReprEnter
func ReprEnter(object *python.PyObject) int {
	return int(C.Py_ReprEnter(toC(object)))
}

// ReprLeave Py_ReprLeave
func ReprLeave(object *python.PyObject) {
	C.Py_ReprLeave(toC(object))
}
