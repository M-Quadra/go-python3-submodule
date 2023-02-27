package pyerr

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

// SetString PyErr_SetString
func SetString(typePy *python.PyObject, message string) {
	messageC := C.CString(message)
	defer C.free(unsafe.Pointer(messageC))

	C.PyErr_SetString(toC(typePy), messageC)
}

// SetObject PyErr_SetObject
func SetObject(typePy, value *python.PyObject) {
	C.PyErr_SetObject(toC(typePy), toC(value))
}

// PyErr_Format

// PyErr_FormatV

// SetNone PyErr_SetNone
func SetNone(typePy *python.PyObject) {
	C.PyErr_SetNone(toC(typePy))
}

// BadArgument PyErr_BadArgument
func BadArgument() int {
	return int(C.PyErr_BadArgument())
}

// NoMemory PyErr_NoMemory
func NoMemory() *python.PyObject {
	return toObject(C.PyErr_NoMemory())
}

// PyErr_SetFromErrno

// PyErr_SetFromErrnoWithFilenameObject

// PyErr_SetFromErrnoWithFilenameObjects

// PyErr_SetFromErrnoWithFilename

// PyErr_SetFromWindowsErr

// PyErr_SetExcFromWindowsErr

// PyErr_SetFromWindowsErrWithFilename

// PyErr_SetExcFromWindowsErrWithFilenameObject

// PyErr_SetExcFromWindowsErrWithFilenameObjects

// PyErr_SetExcFromWindowsErrWithFilename

// SetImportError PyErr_SetImportError
func SetImportError(msg, name, path *python.PyObject) *python.PyObject {
	return toObject(C.PyErr_SetImportError(toC(msg), toC(name), toC(path)))
}

// SyntaxLocationObject PyErr_SyntaxLocationObject
func SyntaxLocationObject(filename *python.PyObject, lineno, colOffset int) {
	filenameC := toC(filename)
	if filenameC == nil || Occurred() == nil {
		return
	}

	C.PyErr_SyntaxLocationObject(filenameC, (C.int)(lineno), (C.int)(colOffset))
}

// SyntaxLocationEx PyErr_SyntaxLocationEx
func SyntaxLocationEx(filename string, lineno, colOffset int) {
	if Occurred() == nil {
		return
	}

	filenameC := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))

	C.PyErr_SyntaxLocationEx(filenameC, (C.int)(lineno), (C.int)(colOffset))
}

// SyntaxLocation PyErr_SyntaxLocation
func SyntaxLocation(filename string, lineno int) {
	if Occurred() == nil {
		return
	}

	filenameC := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))

	C.PyErr_SyntaxLocation(filenameC, (C.int)(lineno))
}

// BadInternalCall PyErr_BadInternalCall
func BadInternalCall() {
	C.PyErr_BadInternalCall()
}
