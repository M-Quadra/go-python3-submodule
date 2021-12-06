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

	python "github.com/M-Quadra/go-python3-submodule/v9"
)

// Call PyObject_Call
func Call(callable, args, kwargs *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_Call(toC(callable), toC(args), toC(kwargs)))
}

// CallNoArgs PyObject_CallNoArgs
func CallNoArgs(callable *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_CallNoArgs(toC(callable)))
}

// CallOneArg PyObject_CallOneArg
func CallOneArg(callable, arg *python.PyObject) *python.PyObject {
	if arg == nil {
		return CallNoArgs(callable)
	}

	return toObject(C.PyObject_CallOneArg(toC(callable), toC(arg)))
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
		return CallNoArgs(callable)
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
		return CallMethodNoArgs(obj, name)
	}

	argsC := make([]*C.PyObject, 0, len(args))
	for _, v := range args {
		argsC = append(argsC, toC(v))
	}

	return toObject(C.cgo_PyObject_CallMethodObjArgs(toC(obj), toC(name), (C.int)(len(args)), (**C.PyObject)(unsafe.Pointer(&argsC[0]))))
}

// CallMethodNoArgs PyObject_CallMethodNoArgs
func CallMethodNoArgs(obj, name *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_CallMethodNoArgs(toC(obj), toC(name)))
}

// CallMethodOneArg PyObject_CallMethodOneArg
func CallMethodOneArg(obj, name, arg *python.PyObject) *python.PyObject {
	return toObject(C.PyObject_CallMethodOneArg(toC(obj), toC(name), toC(arg)))
}

// PyObject_Vectorcall

// PyObject_VectorcallDict

// PyObject_VectorcallMethod
