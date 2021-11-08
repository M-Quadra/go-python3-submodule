package python

import (
	/*
		#cgo pkg-config: python3-embed
		#include "Python.h"
	*/
	"C"
)

// PyObject C.PyObject
type PyObject C.PyObject

// PyThreadState C.PyThreadState
type PyThreadState C.PyThreadState

// PyGILState C.PyGILState_STATE
type PyGILState C.PyGILState_STATE
