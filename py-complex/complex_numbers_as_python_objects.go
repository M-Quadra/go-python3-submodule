package pycomplex

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import python "github.com/M-Quadra/go-python3-submodule"

// Check PyComplex_Check
func Check(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyComplex_Check(toC(p)) != 0
}

// CheckExact PyComplex_CheckExact
func CheckExact(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyComplex_CheckExact(toC(p)) != 0
}

// PyComplex_FromCComplex

// FromFloat64s PyComplex_FromDoubles
func FromFloat64s(real, imag float64) *python.PyObject {
	return toObject(C.PyComplex_FromDoubles((C.double)(real), (C.double)(imag)))
}

// RealAsFloat64 PyComplex_RealAsDouble
func RealAsFloat64(op *python.PyObject) float64 {
	if op == nil {
		return 0
	}

	return float64(C.PyComplex_RealAsDouble(toC(op)))
}

// ImagAsFloat64 PyComplex_ImagAsDouble
func ImagAsFloat64(op *python.PyObject) float64 {
	if op == nil {
		return 0
	}

	return float64(C.PyComplex_ImagAsDouble(toC(op)))
}

// PyComplex_AsCComplex
