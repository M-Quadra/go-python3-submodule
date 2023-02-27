package pylong

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

// Type PyLong_Type
var Type = (*python.PyObject)(unsafe.Pointer(&C.PyLong_Type))

// Check PyLong_Check
func Check(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyLong_Check(toC(p)) != 0
}

// CheckExact PyLong_CheckExact
func CheckExact(p *python.PyObject) bool {
	if p == nil {
		return false
	}

	return C.cgo_PyLong_CheckExact(toC(p)) != 0
}

// FromInt PyLong_FromLong
func FromInt(v int) *python.PyObject {
	return toObject(C.PyLong_FromLong(C.long(v)))
}

// FromUint PyLong_FromUnsignedLong
func FromUint(v uint) *python.PyObject {
	return toObject(C.PyLong_FromUnsignedLong(C.ulong(v)))
}

// PyLong_FromSsize_t

// PyLong_FromSize_t

// FromInt64 PyLong_FromLongLong
func FromInt64(v int64) *python.PyObject {
	return toObject(C.PyLong_FromLongLong(C.longlong(v)))
}

// FromUint64 PyLong_FromUnsignedLongLong
func FromUint64(v uint64) *python.PyObject {
	return toObject(C.PyLong_FromUnsignedLongLong(C.ulonglong(v)))
}

// FromFloat64 PyLong_FromDouble
func FromFloat64(v float64) *python.PyObject {
	return toObject(C.PyLong_FromDouble(C.double(v)))
}

// FromString PyLong_FromString
func FromString(str string, base int) *python.PyObject {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return toObject(C.PyLong_FromString(cstr, nil, (C.int)(base)))
}

// PyLong_FromUnicode

// FromUnicodeObject PyLong_FromUnicodeObject
func FromUnicodeObject(u *python.PyObject, base int) *python.PyObject {
	if u == nil {
		return nil
	}

	return toObject(C.PyLong_FromUnicodeObject(toC(u), (C.int)(base)))
}

// PyLong_FromVoidPtr

// AsInt PyLong_AsLong
func AsInt(obj *python.PyObject) int {
	return int(C.PyLong_AsLong(toC(obj)))
}

// AsIntAndOverflow PyLong_AsLongAndOverflow
func AsIntAndOverflow(obj *python.PyObject) (int, int) {
	overflow := (C.int)(0)
	v := C.PyLong_AsLongAndOverflow(toC(obj), &overflow)
	return int(v), int(overflow)
}

// AsInt64 PyLong_AsLongLong
func AsInt64(obj *python.PyObject) int64 {
	return int64(C.PyLong_AsLongLong(toC(obj)))
}

// AsInt64AndOverflow PyLong_AsLongLongAndOverflow
func AsInt64AndOverflow(obj *python.PyObject) (int64, int) {
	overflow := (C.int)(0)
	v := C.PyLong_AsLongLongAndOverflow(toC(obj), &overflow)
	return int64(v), int(overflow)
}

// PyLong_AsSsize_t

// AsUint PyLong_AsUnsignedLong
func AsUint(pylong *python.PyObject) uint {
	return uint(C.PyLong_AsUnsignedLong(toC(pylong)))
}

// PyLong_AsSize_t

// AsUint64 PyLong_AsUnsignedLongLong
func AsUint64(pylong *python.PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLongLong(toC(pylong)))
}

// AsUintMask PyLong_AsUnsignedLongMask
func AsUintMask(obj *python.PyObject) uint {
	return uint(C.PyLong_AsUnsignedLongMask(toC(obj)))
}

// AsUint64Mask PyLong_AsUnsignedLongLongMask
func AsUint64Mask(obj *python.PyObject) uint64 {
	return uint64(C.PyLong_AsUnsignedLongLongMask(toC(obj)))
}

// AsFloat64 PyLong_AsDouble
func AsFloat64(pylong *python.PyObject) float64 {
	return float64(C.PyLong_AsDouble(toC(pylong)))
}

// PyLong_AsVoidPtr
