package py

import (
	/*
		#include "Python.h"
	*/
	"C"
)

var (
	// False Py_False
	False = toObject(C.Py_False)
	// True Py_True
	True = toObject(C.Py_True)
)
