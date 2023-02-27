package pysys

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

// GetObject PySys_GetObject
func GetObject(name string) *python.PyObject {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	return toObject(C.PySys_GetObject(nameC))
}

// SetObject PySys_SetObject
func SetObject(name string, v *python.PyObject) bool {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	return C.PySys_SetObject(nameC, toC(v)) == 0
}

// ResetWarnOptions PySys_ResetWarnOptions
func ResetWarnOptions() {
	C.PySys_ResetWarnOptions()
}

// AddWarnOption PySys_AddWarnOption
func AddWarnOption(s string) {
	sC := C.CString(s)
	defer C.free(unsafe.Pointer(sC))

	ws := C.Py_DecodeLocale(sC, nil)
	if ws == nil {
		return
	}
	defer C.PyMem_RawFree(unsafe.Pointer(ws))

	C.PySys_AddWarnOption(ws)
}

// AddWarnOptionUnicode PySys_AddWarnOptionUnicode
func AddWarnOptionUnicode(unicode *python.PyObject) {
	C.PySys_AddWarnOptionUnicode(toC(unicode))
}

// SetPath PySys_SetPath
func SetPath(path string) {
	pathC := C.CString(path)
	defer C.free(unsafe.Pointer(pathC))

	wpath := C.Py_DecodeLocale(pathC, nil)
	if wpath == nil {
		return
	}
	defer C.PyMem_RawFree(unsafe.Pointer(wpath))

	C.PySys_SetPath(wpath)
}

// PySys_WriteStdout

// PySys_WriteStderr

// PySys_FormatStdout

// PySys_FormatStderr

// AddXOption PySys_AddXOption
func AddXOption(s string) {
	sC := C.CString(s)
	defer C.free(unsafe.Pointer(sC))

	ws := C.Py_DecodeLocale(sC, nil)
	if ws == nil {
		return
	}
	defer C.PyMem_RawFree(unsafe.Pointer(ws))

	C.PySys_AddXOption(ws)
}

// GetXOptions PySys_GetXOptions
func GetXOptions() *python.PyObject {
	return toObject(C.PySys_GetXOptions())
}

// PySys_Audit

// PySys_AddAuditHook
