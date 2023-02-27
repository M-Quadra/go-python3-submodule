package pyunicode

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule/v11"
)

// New PyUnicode_New
func New(size int, maxchar rune) *python.PyObject {
	return toObject(C.PyUnicode_New(C.Py_ssize_t(size), C.Py_UCS4(maxchar)))
}

// PyUnicode_FromKindAndData

// PyUnicode_FromStringAndSize

// FromString PyUnicode_FromString
func FromString(u string) *python.PyObject {
	uC := C.CString(u)
	defer C.free(unsafe.Pointer(uC))

	return toObject(C.PyUnicode_FromString(uC))
}

// PyUnicode_FromFormat

// PyUnicode_FromFormatV

// FromEncodedObject PyUnicode_FromEncodedObject
func FromEncodedObject(obj *python.PyObject, encoding string, errors string) *python.PyObject {
	encodingC := C.CString(encoding)
	defer C.free(unsafe.Pointer(encodingC))

	errorsC := C.CString(errors)
	defer C.free(unsafe.Pointer(errorsC))

	return toObject(C.PyUnicode_FromEncodedObject(toC(obj), encodingC, errorsC))
}

// GetLength PyUnicode_GetLength
func GetLength(unicode *python.PyObject) int {
	if unicode == nil {
		return 0
	}

	return int(C.PyUnicode_GetLength(toC(unicode)))
}

// CopyCharacters PyUnicode_CopyCharacters
func CopyCharacters(to *python.PyObject, toStart int, from *python.PyObject, fromStart, howMany int) int {
	if to == nil || from == nil {
		return -1
	}

	return int(C.PyUnicode_CopyCharacters(toC(to), C.Py_ssize_t(toStart), toC(from), C.Py_ssize_t(fromStart), C.Py_ssize_t(howMany)))
}

// Fill PyUnicode_Fill
func Fill(unicode *python.PyObject, start, length int, fillChar rune) int {
	if unicode == nil {
		return 0
	}

	return int(C.PyUnicode_Fill(toC(unicode), C.Py_ssize_t(start), C.Py_ssize_t(length), C.Py_UCS4(fillChar)))
}

// WriteChar PyUnicode_WriteChar
func WriteChar(unicode *python.PyObject, index int, character rune) int {
	if unicode == nil {
		return 0
	}

	return int(C.PyUnicode_WriteChar(toC(unicode), C.Py_ssize_t(index), C.Py_UCS4(character)))
}

// ReadChar PyUnicode_ReadChar
func ReadChar(unicode *python.PyObject, index int) rune {
	if unicode == nil {
		return 0
	}

	return rune(C.PyUnicode_ReadChar(toC(unicode), C.Py_ssize_t(index)))
}

// Substring PyUnicode_Substring
func Substring(str *python.PyObject, start, end int) *python.PyObject {
	if str == nil {
		return nil
	}

	return toObject(C.PyUnicode_Substring(toC(str), C.Py_ssize_t(start), C.Py_ssize_t(end)))
}

// PyUnicode_AsUCS4

// PyUnicode_AsUCS4Copy
