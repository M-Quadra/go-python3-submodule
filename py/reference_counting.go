package py

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import "python"

// Py_INCREF

// Py_XINCREF

// DecRef Py_DECREF
func DecRef(o *python.PyObject) {
	C.Py_DecRef(toC(o))
}

// Py_XDECREF

// Py_CLEAR
