package pyunicode

import (
	/*
		#include "Python.h"
	*/
	"C"
)
import "python"

// AsString PyUnicode_AsUTF8
func AsString(unicode *python.PyObject) string {
	if unicode == nil {
		return ""
	}

	return C.GoString(C.PyUnicode_AsUTF8(toC(unicode)))
}
