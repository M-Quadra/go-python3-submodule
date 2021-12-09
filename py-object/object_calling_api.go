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

	python "github.com/M-Quadra/go-python3-submodule/v8"
)

// Call PyObject_Call
func Call(callable, args, kwargs *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_Call(toC(callable), toC(args), toC(kwargs)))
}

// CallObject PyObject_CallObject
func CallObject(callable, args *python.PyObject) *python.PyObject {
	if callable == nil {
		return nil
	}

	return toObject(C.PyObject_CallObject(toC(callable), toC(args)))
}

// PyObject_CallFunction

// PyObject_CallMethod

// MaxVariadicLength is the maximum number of arguments that can be passed to a variadic C function due to a cgo limitation
const MaxVariadicLength = 20

// CallFunctionObjArgs PyObject_CallFunctionObjArgs
func CallFunctionObjArgs(callable *python.PyObject, args ...*python.PyObject) *python.PyObject {
	if len(args) > MaxVariadicLength {
		// CallFunctionObjArgs: too many arrguments
		return nil
	}
	if len(args) <= 0 {
		toObject(C.cgo_PyObject_CallFunctionObjArgs(toC(callable), 0, nil))
	}

	argsC := make([]*C.PyObject, 0, len(args))
	for _, v := range args {
		argsC = append(argsC, toC(v))
	}

	return toObject(C.cgo_PyObject_CallFunctionObjArgs(toC(callable), (C.int)(len(args)), (**C.PyObject)(unsafe.Pointer(&argsC[0]))))
}

// CallMethodObjArgs PyObject_CallMethodObjArgs
func CallMethodObjArgs(obj, name *python.PyObject, args ...*python.PyObject) *python.PyObject {
	if len(args) > MaxVariadicLength {
		// CallMethodObjArgs: too many arrguments
		return nil
	}
	if len(args) <= 0 {
		return toObject(C.cgo_PyObject_CallMethodObjArgs(toC(obj), toC(name), 0, nil))
	}

	argsC := make([]*C.PyObject, 0, len(args))
	for _, v := range args {
		argsC = append(argsC, toC(v))
	}

	return toObject(C.cgo_PyObject_CallMethodObjArgs(toC(obj), toC(name), (C.int)(len(args)), (**C.PyObject)(unsafe.Pointer(&argsC[0]))))
}

// PyObject_Vectorcall

// PyObject_VectorcallDict

// PyObject_VectorcallMethod
