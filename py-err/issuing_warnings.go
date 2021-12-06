package pyerr

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule/v9"
)

// WarnEx PyErr_WarnEx
func WarnEx(category *python.PyObject, message string, stackLevel int) bool {
	messageC := C.CString(message)
	defer C.free(unsafe.Pointer(messageC))

	return C.PyErr_WarnEx(toC(category), messageC, (C.Py_ssize_t)(stackLevel)) == 0
}

// SetImportErrorSubclass PyErr_SetImportErrorSubclass
func SetImportErrorSubclass(exception, msg, name, path *python.PyObject) *python.PyObject {
	if exception == nil {
		return nil
	}

	return toObject(C.PyErr_SetImportErrorSubclass(toC(exception), toC(msg), toC(name), toC(path)))
}

// WarnExplicitObject PyErr_WarnExplicitObject
func WarnExplicitObject(category, message, filename *python.PyObject, lineno int, module, registry *python.PyObject) bool {
	if message == nil {
		return false
	}

	return C.PyErr_WarnExplicitObject(toC(category), toC(message), toC(filename), (C.int)(lineno), toC(module), toC(registry)) == 0
}

// WarnExplicit PyErr_WarnExplicit
func WarnExplicit(category *python.PyObject, message, filename string, lineno int, module string, registry *python.PyObject) bool {
	messageC := C.CString(message)
	defer C.free(unsafe.Pointer(messageC))
	filenameC := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))
	moduleC := C.CString(module)
	defer C.free(unsafe.Pointer(moduleC))

	return C.PyErr_WarnExplicit(toC(category), messageC, filenameC, (C.int)(lineno), moduleC, toC(registry)) == 0
}

// PyErr_WarnFormat

// PyErr_ResourceWarning
