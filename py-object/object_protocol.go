package pyobject

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

// PyObject_Print

// HasAttr PyObject_HasAttr
func HasAttr(o, attrName *python.PyObject) bool {
	if o == nil || attrName == nil {
		return false
	}

	return C.PyObject_HasAttr(toC(o), toC(attrName)) == 1
}

// HasAttrString PyObject_HasAttrString
func HasAttrString(o *python.PyObject, attrName string) bool {
	if o == nil {
		return false
	}

	attrNameC := C.CString(attrName)
	defer C.free(unsafe.Pointer(attrNameC))

	return C.PyObject_HasAttrString(toC(o), attrNameC) == 1
}

// GetAttr PyObject_GetAttr
func GetAttr(o, attrName *python.PyObject) *python.PyObject {
	if o == nil || attrName == nil {
		return nil
	}

	return toObject(C.PyObject_GetAttr(toC(o), toC(attrName)))
}

// GetAttrString PyObject_GetAttrString
func GetAttrString(o *python.PyObject, attrName string) *python.PyObject {
	if o == nil || len(attrName) <= 0 {
		return nil
	}

	attrNameC := C.CString(attrName)
	defer C.free(unsafe.Pointer(attrNameC))

	return toObject(C.PyObject_GetAttrString(toC(o), attrNameC))
}

// GenericGetAttr PyObject_GenericGetAttr
func GenericGetAttr(o, name *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_GenericGetAttr(toC(o), toC(name)))
}

// SetAttr PyObject_SetAttr
func SetAttr(o, attrName, v *python.PyObject) bool {
	if o == nil || attrName == nil {
		return false
	}

	return C.PyObject_SetAttr(toC(o), toC(attrName), toC(v)) == 0
}

// SetAttrString PyObject_SetAttrString
func SetAttrString(o *python.PyObject, attrName string, v *python.PyObject) bool {
	if o == nil {
		return false
	}

	attrNameC := C.CString(attrName)
	defer C.free(unsafe.Pointer(attrNameC))

	return C.PyObject_SetAttrString(toC(o), attrNameC, toC(v)) == 0
}

// GenericSetAttr PyObject_GenericSetAttr
func GenericSetAttr(o, name, value *python.PyObject) bool {
	return C.PyObject_GenericSetAttr(toC(o), toC(name), toC(value)) == 0
}

// DelAttr PyObject_DelAttr
func DelAttr(o, attrName *python.PyObject) bool {
	if o == nil || attrName == nil {
		return false
	}

	return C.cgo_PyObject_DelAttr(toC(o), toC(attrName)) != -1
}

// DelAttrString PyObject_DelAttrString
func DelAttrString(o *python.PyObject, attrName string) bool {
	if o == nil {
		return false
	}

	attrNameC := C.CString(attrName)
	defer C.free(unsafe.Pointer(attrNameC))

	return C.cgo_PyObject_DelAttrString(toC(o), attrNameC) != -1
}

// PyObject_GenericGetDict

// PyObject_GenericSetDict

// RichCompare PyObject_RichCompare
func RichCompare(o1, o2 *python.PyObject, opid int) *python.PyObject {
	if opid < 0 || 5 < opid {
		return nil
	}

	return toObject(C.PyObject_RichCompare(toC(o1), toC(o2), (C.int)(opid)))
}

// RichCompareBool PyObject_RichCompareBool
func RichCompareBool(o1, o2 *python.PyObject, opid int) int {
	if opid < 0 || 5 < opid {
		return -1
	}

	return int(C.PyObject_RichCompareBool(toC(o1), toC(o2), (C.int)(opid)))
}

// Repr PyObject_Repr
func Repr(o *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_Repr(toC(o)))
}

// ASCII PyObject_ASCII
func ASCII(o *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_ASCII(toC(o)))
}

// Str PyObject_Str
func Str(o *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_Str(toC(o)))
}

// Bytes PyObject_Bytes
func Bytes(o *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_Bytes(toC(o)))
}

// IsSubclass PyObject_IsSubclass
func IsSubclass(derived, cls *python.PyObject) int {
	if derived == nil || cls == nil {
		return -1
	}

	return int(C.PyObject_IsSubclass(toC(derived), toC(cls)))
}

// IsInstance PyObject_IsInstance
func IsInstance(inst, cls *python.PyObject) int {
	if inst == nil || cls == nil {
		return -1
	}

	return int(C.PyObject_IsInstance(toC(inst), toC(cls)))
}

// Hash PyObject_Hash
func Hash(o *python.PyObject) int {
	if o == nil {
		return -1
	}

	return int(C.PyObject_Hash(toC(o)))
}

// HashNotImplemented PyObject_HashNotImplemented
func HashNotImplemented(o *python.PyObject) int {
	if o == nil {
		return -1
	}

	return int(C.PyObject_HashNotImplemented(toC(o)))
}

// IsTrue PyObject_IsTrue
func IsTrue(o *python.PyObject) int {
	if o == nil {
		return -1
	}

	return int(C.PyObject_IsTrue(toC(o)))
}

// Not PyObject_Not
func Not(o *python.PyObject) int {
	if o == nil {
		return -1
	}

	return int(C.PyObject_Not(toC(o)))
}

// Type PyObject_Type
func Type(o *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_Type(toC(o)))
}

// PyObject_TypeCheck

// Size PyObject_Size
func Size(o *python.PyObject) int {
	return int(C.PyObject_Size(toC(o)))
}

// Length PyObject_Length
func Length(o *python.PyObject) int {
	return int(C.PyObject_Length(toC(o)))
}

// LengthHint PyObject_LengthHint
func LengthHint(o *python.PyObject, defaultvalue int) int {
	if o == nil {
		return defaultvalue
	}

	return int(C.PyObject_LengthHint(toC(o), (C.Py_ssize_t)(defaultvalue)))
}

// GetItem PyObject_GetItem
func GetItem(o, key *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_GetItem(toC(o), toC(key)))
}

// SetItem PyObject_SetItem
func SetItem(o, key, val *python.PyObject) bool {
	return C.PyObject_SetItem(toC(o), toC(key), toC(val)) == 0
}

// DelItem PyObject_DelItem
func DelItem(o, key *python.PyObject) bool {
	return C.PyObject_DelItem(toC(o), toC(key)) != -1
}

// Dir PyObject_Dir
func Dir(o *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_Dir(toC(o)))
}

// GetIter PyObject_GetIter
func GetIter(o *python.PyObject) *python.PyObject {
	if o == nil {
		return nil
	}

	return toObject(C.PyObject_GetIter(toC(o)))
}
