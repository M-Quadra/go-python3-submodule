package py

import (
	/*
		#include "Python.h"
	*/
	"C"
)

// None Py_None
var None = toObject(C.Py_None)
