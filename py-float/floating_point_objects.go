package pyfloat

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import (
	python "github.com/M-Quadra/go-python3-submodule/v11"
	"github.com/M-Quadra/go-python3-submodule/v11/py"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v11/py-unicode"
)

// Check PyFloat_Check
func Check(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyFloat_Check(toC(p)) != 0
}

// CheckExact PyFloat_CheckExact
func CheckExact(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyFloat_CheckExact(toC(p)) != 0
}

// FromStringPy PyFloat_FromString
func FromStringPy(str *python.PyObject) *python.PyObject {
	if str == nil {
		return nil
	}

	return toObject(C.PyFloat_FromString(toC(str)))
}

// FromString PyFloat_FromString
func FromString(str string) *python.PyObject {
	strPy := pyunicode.FromString(str)
	defer py.DecRef(strPy)

	return FromStringPy(strPy)
}

// FromFloat64 PyFloat_FromDouble
func FromFloat64(v float64) *python.PyObject {
	return toObject(C.PyFloat_FromDouble((C.double)(v)))
}

// AsFloat64 PyFloat_AsDouble
func AsFloat64(pyfloat *python.PyObject) float64 {
	return float64(C.PyFloat_AsDouble(toC(pyfloat)))
}

// PyFloat_AS_DOUBLE

// GetInfo PyFloat_GetInfo
func GetInfo() *python.PyObject {
	return toObject(C.PyFloat_GetInfo())
}

// GetMax PyFloat_GetMax
func GetMax() float64 {
	return float64(C.PyFloat_GetMax())
}

// GetMin PyFloat_GetMin
func GetMin() float64 {
	return float64(C.PyFloat_GetMin())
}
