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
