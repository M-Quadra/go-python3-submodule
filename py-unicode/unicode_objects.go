package pyunicode

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import python "github.com/M-Quadra/go-python3-submodule/v9"

// Check PyUnicode_Check
func Check(o *python.PyObject) bool {
	if o == nil {
		return false
	}

	return C.cgo_PyUnicode_Check(toC(o)) != 0
}

// CheckExact PyUnicode_CheckExact
func CheckExact(o *python.PyObject) bool {
	if o == nil {
		return false
	}

	return C.cgo_PyUnicode_CheckExact(toC(o)) != 0
}

// PyUnicode_READY

// PyUnicode_GET_LENGTH

// PyUnicode_1BYTE_DATA
// PyUnicode_2BYTE_DATA
// PyUnicode_4BYTE_DATA

// PyUnicode_KIND

// PyUnicode_DATA

// PyUnicode_WRITE

// PyUnicode_READ

// PyUnicode_READ_CHAR

// PyUnicode_GET_SIZE

// PyUnicode_GET_DATA_SIZE

// PyUnicode_AS_UNICODE
// PyUnicode_AS_DATA

// PyUnicode_IsIdentifier
