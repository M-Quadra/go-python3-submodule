package pyimport

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule/v10"
)

// ImportModule PyImport_ImportModule
func ImportModule(name string) *python.PyObject {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	return toObject(C.PyImport_ImportModule(nameC))
}

// PyImport_ImportModuleNoBlock

// ImportModuleEx PyImport_ImportModuleEx
func ImportModuleEx(name string, globals, locals, fromlist *python.PyObject) *python.PyObject {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	return toObject(C.cgo_PyImport_ImportModuleEx(nameC, toC(globals), toC(locals), toC(fromlist)))
}

// ImportModuleLevelObject PyImport_ImportModuleLevelObject
func ImportModuleLevelObject(name, globals, locals, fromlist *python.PyObject, level int) *python.PyObject {
	return toObject(C.PyImport_ImportModuleLevelObject(toC(name), toC(globals), toC(locals), toC(fromlist), (C.int)(level)))
}

// ImportModuleLevel PyImport_ImportModuleLevel
func ImportModuleLevel(name string, globals, locals, fromlist *python.PyObject, level int) *python.PyObject {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	return toObject(C.PyImport_ImportModuleLevel(nameC, toC(globals), toC(locals), toC(fromlist), (C.int)(level)))
}

// Import PyImport_Import
func Import(name *python.PyObject) *python.PyObject {
	return toObject(C.PyImport_Import(toC(name)))
}

// ReloadModule PyImport_ReloadModule
func ReloadModule(m *python.PyObject) *python.PyObject {
	if m == nil {
		return nil
	}

	return toObject(C.PyImport_ReloadModule(toC(m)))
}

// AddModuleObject PyImport_AddModuleObject
func AddModuleObject(name *python.PyObject) *python.PyObject {
	if name == nil {
		return nil
	}

	return toObject(C.PyImport_AddModuleObject(toC(name)))
}

// AddModule PyImport_AddModule
func AddModule(name string) *python.PyObject {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	return toObject(C.PyImport_AddModule(nameC))
}

// ExecCodeModule PyImport_ExecCodeModule
func ExecCodeModule(name string, co *python.PyObject) *python.PyObject {
	if co == nil {
		return nil
	}

	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	return toObject(C.PyImport_ExecCodeModule(nameC, toC(co)))
}

// ExecCodeModuleEx PyImport_ExecCodeModuleEx
func ExecCodeModuleEx(name string, co *python.PyObject, pathname string) *python.PyObject {
	if co == nil {
		return nil
	}

	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	pathnameC := C.CString(pathname)
	defer C.free(unsafe.Pointer(pathnameC))

	return toObject(C.PyImport_ExecCodeModuleEx(nameC, toC(co), pathnameC))
}

// ExecCodeModuleObject PyImport_ExecCodeModuleObject
func ExecCodeModuleObject(name, co, pathname, cpathname *python.PyObject) *python.PyObject {
	if name == nil || co == nil {
		return nil
	}

	return toObject(C.PyImport_ExecCodeModuleObject(toC(name), toC(co), toC(pathname), toC(cpathname)))
}

// ExecCodeModuleWithPathnames PyImport_ExecCodeModuleWithPathnames
func ExecCodeModuleWithPathnames(name string, co *python.PyObject, pathname, cpathname string) *python.PyObject {
	if co == nil {
		return nil
	}

	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	pathnameC := C.CString(pathname)
	defer C.free(unsafe.Pointer(pathnameC))
	cpathnameC := C.CString(cpathname)
	defer C.free(unsafe.Pointer(cpathnameC))

	return toObject(C.PyImport_ExecCodeModuleWithPathnames(nameC, toC(co), pathnameC, cpathnameC))
}

// GetMagicNumber PyImport_GetMagicNumber
func GetMagicNumber() int {
	return (int)(C.PyImport_GetMagicNumber())
}

// GetMagicTag PyImport_GetMagicTag
func GetMagicTag() string {
	return C.GoString(C.PyImport_GetMagicTag())
}

// GetModuleDict PyImport_GetModuleDict
func GetModuleDict() *python.PyObject {
	return toObject(C.PyImport_GetModuleDict())
}

// GetModule PyImport_GetModule
func GetModule(name *python.PyObject) *python.PyObject {
	if name == nil {
		return nil
	}

	return toObject(C.PyImport_GetModule(toC(name)))
}

// GetImporter PyImport_GetImporter
func GetImporter(path *python.PyObject) *python.PyObject {
	if path == nil {
		return nil
	}

	return toObject(C.PyImport_GetImporter(toC(path)))
}

// ImportFrozenModuleObject PyImport_ImportFrozenModuleObject
func ImportFrozenModuleObject(name *python.PyObject) int {
	return (int)(C.PyImport_ImportFrozenModuleObject(toC(name)))
}

// ImportFrozenModule PyImport_ImportFrozenModule
func ImportFrozenModule(name string) int {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	return (int)(C.PyImport_ImportFrozenModule(nameC))
}

// PyImport_AppendInittab

// PyImport_ExtendInittab
