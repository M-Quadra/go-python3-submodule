package pyeval

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule/v8"
)

// InitThreads PyEval_InitThreads
func InitThreads() {
	C.PyEval_InitThreads()
}

// ThreadsInitialized PyEval_ThreadsInitialized
func ThreadsInitialized() bool {
	return C.PyEval_ThreadsInitialized() != 0
}

// SaveThread PyEval_SaveThread
func SaveThread() *python.PyThreadState {
	return (*python.PyThreadState)(unsafe.Pointer(C.PyEval_SaveThread()))
}

// RestoreThread PyEval_RestoreThread
func RestoreThread(tstate *python.PyThreadState) {
	C.PyEval_RestoreThread((*C.PyThreadState)(unsafe.Pointer(tstate)))
}
