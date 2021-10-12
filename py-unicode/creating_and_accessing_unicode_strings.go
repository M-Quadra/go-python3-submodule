package pyunicode

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import (
	"unsafe"

	python "github.com/M-Quadra/go-python3-submodule"
)

// PyUnicode_New

// PyUnicode_FromKindAndData

// PyUnicode_FromStringAndSize

// FromString PyUnicode_FromString
func FromString(u string) *python.PyObject {
	uC := C.CString(u)
	defer C.free(unsafe.Pointer(uC))

	return toObject(C.PyUnicode_FromString(uC))
}

// PyUnicode_FromFormat

// PyUnicode_FromFormatV

// PyUnicode_FromEncodedObject

// PyUnicode_GetLength

// PyUnicode_CopyCharacters

// PyUnicode_Fill

// PyUnicode_WriteChar

// PyUnicode_ReadChar

// PyUnicode_Substring

// PyUnicode_AsUCS4

// PyUnicode_AsUCS4Copy
