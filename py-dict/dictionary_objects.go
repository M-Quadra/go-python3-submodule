package pydict

import (
	/*
		#include "Python.h"
		#include "cgo_bridge.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule/v11"
)

// Check PyDict_Check
func Check(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyDict_Check(toC(p)) != 0
}

// CheckExact PyDict_CheckExact
func CheckExact(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyDict_CheckExact(toC(p)) != 0
}

// New PyDict_New
func New() *python.PyObject {
	return toObject(C.PyDict_New())
}

// ProxyNew PyDictProxy_New
func ProxyNew(mapping *python.PyObject) *python.PyObject {
	if mapping == nil {
		return nil
	}

	return toObject(C.PyDictProxy_New(toC(mapping)))
}

// Clear PyDict_Clear
func Clear(p *python.PyObject) {
	if p == nil {
		return
	}

	C.PyDict_Clear(toC(p))
}

// Contains PyDict_Contains
func Contains(p, key *python.PyObject) int {
	if key == nil {
		return 0
	}

	return int(C.PyDict_Contains(toC(p), toC(key)))
}

// Copy PyDict_Copy
func Copy(p *python.PyObject) *python.PyObject {
	return toObject(C.PyDict_Copy(toC(p)))
}

// SetItem PyDict_SetItem
func SetItem(p, key, val *python.PyObject) bool {
	if p == nil || key == nil || val == nil {
		return false
	}

	return C.PyDict_SetItem(toC(p), toC(key), toC(val)) == 0
}

// SetItemString PyDict_SetItemString
func SetItemString(p *python.PyObject, key string, val *python.PyObject) bool {
	if val == nil {
		return false
	}

	keyC := C.CString(key)
	defer C.free(unsafe.Pointer(keyC))

	return C.PyDict_SetItemString(toC(p), keyC, toC(val)) == 0
}

// DelItem PyDict_DelItem
func DelItem(p, key *python.PyObject) bool {
	if p == nil || key == nil {
		return false
	}

	return C.PyDict_DelItem(toC(p), toC(key)) == 0
}

// DelItemString PyDict_DelItemString
func DelItemString(p *python.PyObject, key string) bool {
	keyC := C.CString(key)
	defer C.free(unsafe.Pointer(keyC))

	return C.PyDict_DelItemString(toC(p), keyC) == 0
}

// GetItem PyDict_GetItem
func GetItem(p, key *python.PyObject) *python.PyObject {
	if p == nil || key == nil {
		return nil
	}

	return toObject(C.PyDict_GetItem(toC(p), toC(key)))
}

// GetItemWithError PyDict_GetItemWithError
func GetItemWithError(p, key *python.PyObject) *python.PyObject {
	if key == nil {
		return nil
	}

	return toObject(C.PyDict_GetItemWithError(toC(p), toC(key)))
}

// GetItemString PyDict_GetItemString
func GetItemString(p *python.PyObject, key string) *python.PyObject {
	keyC := C.CString(key)
	defer C.free(unsafe.Pointer(keyC))

	return toObject(C.PyDict_GetItemString(toC(p), keyC))
}

// SetDefault PyDict_SetDefault
func SetDefault(p, key, defaultobj *python.PyObject) *python.PyObject {
	if p == nil || key == nil || defaultobj == nil {
		return nil
	}

	return toObject(C.PyDict_SetDefault(toC(p), toC(key), toC(defaultobj)))
}

// Items PyDict_Items
func Items(p *python.PyObject) *python.PyObject {
	return toObject(C.PyDict_Items(toC(p)))
}

// Keys PyDict_Keys
func Keys(p *python.PyObject) *python.PyObject {
	return toObject(C.PyDict_Keys(toC(p)))
}

// Values PyDict_Values
func Values(p *python.PyObject) *python.PyObject {
	return toObject(C.PyDict_Values(toC(p)))
}

// Size PyDict_Size
func Size(p *python.PyObject) int {
	return int(C.PyDict_Size(toC(p)))
}

// PyDict_Next

// PyDict_Merge

// PyDict_Update

// PyDict_MergeFromSeq2
