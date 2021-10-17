package py

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import python "github.com/M-Quadra/go-python3-submodule"

// IncRef Py_INCREF
func IncRef(o *python.PyObject) {
	C.Py_IncRef(toC(o))
}

// Py_XINCREF

// DecRef Py_DECREF
func DecRef(o *python.PyObject) {
	C.Py_DecRef(toC(o))
}

// Py_XDECREF

// Py_CLEAR
