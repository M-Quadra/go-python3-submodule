package pymodule

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import (
	"python"
	"unsafe"
)

// Check PyModule_Check
func Check(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyModule_Check(toC(p)) != 0
}

// CheckExact PyModule_CheckExact
func CheckExact(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyModule_CheckExact(toC(p)) != 0
}

// NewObject PyModule_NewObject
func NewObject(name *python.PyObject) *python.PyObject {
	if name == nil {
		return nil
	}

	return toObject(C.PyModule_NewObject(toC(name)))
}

// New PyModule_New
func New(name string) *python.PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return toObject(C.PyModule_New(cname))
}

// GetDict PyModule_GetDict
func GetDict(module *python.PyObject) *python.PyObject {
	if module == nil {
		return nil
	}

	return toObject(C.PyModule_GetDict(toC(module)))
}

// GetNameObject PyModule_GetNameObject
func GetNameObject(module *python.PyObject) *python.PyObject {
	if module == nil {
		return nil
	}

	return toObject(C.PyModule_GetNameObject(toC(module)))
}

// GetName PyModule_GetName
func GetName(module *python.PyObject) string {
	if module == nil {
		return ""
	}

	nameC := C.PyModule_GetName(toC(module))
	return C.GoString(nameC)
}

// GetState PyModule_GetState
func GetState(module *python.PyObject) unsafe.Pointer {
	if module == nil {
		return nil
	}

	return unsafe.Pointer(C.PyModule_GetState(toC(module)))
}

// PyModule_GetDef

// GetFilenameObject PyModule_GetFilenameObject
func GetFilenameObject(module *python.PyObject) *python.PyObject {
	if module == nil {
		return nil
	}

	return toObject(C.PyModule_GetFilenameObject(toC(module)))
}

// GetFilename PyModule_GetFilename
func GetFilename(module *python.PyObject) string {
	if module == nil {
		return ""
	}

	nameC := C.PyModule_GetFilename(toC(module))
	return C.GoString(nameC)
}
