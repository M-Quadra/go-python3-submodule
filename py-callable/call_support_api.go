package pycallable

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import python "github.com/M-Quadra/go-python3-submodule/v9"

// Check PyCallable_Check
func Check(o *python.PyObject) bool {
	if o == nil {
		return false
	}

	return C.PyCallable_Check(toC(o)) != 0
}
