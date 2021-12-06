package pyeval

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import python "github.com/M-Quadra/go-python3-submodule/v9"

// GetBuiltins PyEval_GetBuiltins
func GetBuiltins() *python.PyObject {
	return toObject(C.PyEval_GetBuiltins())
}

// GetLocals PyEval_GetLocals
func GetLocals() *python.PyObject {
	return toObject(C.PyEval_GetLocals())
}

// GetGlobals PyEval_GetGlobals
func GetGlobals() *python.PyObject {
	return toObject(C.PyEval_GetGlobals())
}

// PyEval_GetFrame

// GetFuncName PyEval_GetFuncName
func GetFuncName(funcPy *python.PyObject) string {
	if funcPy == nil {
		return ""
	}

	return C.GoString(C.PyEval_GetFuncName(toC(funcPy)))
}

// GetFuncDesc PyEval_GetFuncDesc
func GetFuncDesc(funcPy *python.PyObject) string {
	return C.GoString(C.PyEval_GetFuncDesc(toC(funcPy)))
}
