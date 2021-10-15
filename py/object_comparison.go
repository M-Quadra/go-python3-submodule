package py

import (
	/*
		#include "Python.h"
	*/
	"C"
)

const (
	// LT Py_LT <
	LT = int(C.Py_LT)
	// LE Py_LE <=
	LE = int(C.Py_LE)
	// EQ Py_EQ ==
	EQ = int(C.Py_EQ)
	// NE Py_NE !=
	NE = int(C.Py_NE)
	// GT Py_GT >
	GT = int(C.Py_GT)
	// GE Py_GE >=
	GE = int(C.Py_GE)
)
