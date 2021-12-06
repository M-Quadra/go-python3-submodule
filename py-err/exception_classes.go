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

// NewException PyErr_NewException
func NewException(name string, base, dict *python.PyObject) *python.PyObject {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	return toObject(C.PyErr_NewException(nameC, toC(base), toC(dict)))
}

// NewExceptionWithDoc PyErr_NewExceptionWithDoc
func NewExceptionWithDoc(name, doc string, base, dict *python.PyObject) *python.PyObject {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	docC := C.CString(doc)
	defer C.free(unsafe.Pointer(docC))

	return toObject(C.PyErr_NewExceptionWithDoc(nameC, docC, toC(base), toC(dict)))
}
