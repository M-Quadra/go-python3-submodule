package pyerr

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import python "github.com/M-Quadra/go-python3-submodule/v8"

// Occurred PyErr_Occurred
func Occurred() *python.PyObject {
	return toObject(C.PyErr_Occurred())
}

// ExceptionMatches PyErr_ExceptionMatches
func ExceptionMatches(exc *python.PyObject) bool {
	return C.PyErr_ExceptionMatches(toC(exc)) != 0
}

// GivenExceptionMatches PyErr_GivenExceptionMatches
func GivenExceptionMatches(given, exc *python.PyObject) bool {
	return C.PyErr_GivenExceptionMatches(toC(given), toC(exc)) != 0
}

// Fetch PyErr_Fetch
func Fetch() (*python.PyObject, *python.PyObject, *python.PyObject) {
	var typeC, valueC, tracebackC *C.PyObject
	C.PyErr_Fetch(&typeC, &valueC, &tracebackC)
	return toObject(typeC), toObject(valueC), toObject(tracebackC)
}

// Restore PyErr_Restore
func Restore(typePy, value, traceback *python.PyObject) {
	C.PyErr_Restore(toC(typePy), toC(value), toC(traceback))
}

// NormalizeException PyErr_NormalizeException
func NormalizeException(exc, val, tb *python.PyObject) (*python.PyObject, *python.PyObject, *python.PyObject) {
	excC, valC, tbC := toC(exc), toC(val), toC(tb)
	C.PyErr_NormalizeException(&excC, &valC, &tbC)
	return toObject(excC), toObject(valC), toObject(tbC)
}

// GetExcInfo PyErr_GetExcInfo
func GetExcInfo() (*python.PyObject, *python.PyObject, *python.PyObject) {
	var typeC, valueC, tracebackC *C.PyObject
	C.PyErr_GetExcInfo(&typeC, &valueC, &tracebackC)
	return toObject(typeC), toObject(valueC), toObject(tracebackC)
}

// SetExcInfo PyErr_SetExcInfo
func SetExcInfo(typePy, value, traceback *python.PyObject) {
	C.PyErr_SetExcInfo(toC(typePy), toC(value), toC(traceback))
}
