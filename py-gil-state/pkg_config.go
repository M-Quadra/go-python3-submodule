package pygilstate

import (
	/*
		#cgo pkg-config: python3-embed
		#include "Python.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule/v11"
)

func toCGILState(state python.PyGILState) C.PyGILState_STATE {
	return *(*C.PyGILState_STATE)(unsafe.Pointer(&state))
}

func toGILState(state C.PyGILState_STATE) python.PyGILState {
	return *(*python.PyGILState)(unsafe.Pointer(&state))
}
