package pysys

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import "unsafe"

// SetArgvEx PySys_SetArgvEx
func SetArgvEx(updatepath bool, argv ...string) {
	if len(argv) <= 0 {
		if updatepath {
			C.PySys_SetArgvEx(0, nil, 1)
		} else {
			C.PySys_SetArgvEx(0, nil, 0)
		}
		return
	}

	argvC := make([]*C.wchar_t, 0, len(argv))
	for _, v := range argv {
		argC := C.CString(v)
		defer C.free(unsafe.Pointer(argC))

		warg := C.Py_DecodeLocale(argC, nil)
		if warg == nil {
			return
		}
		defer C.PyMem_RawFree(unsafe.Pointer(warg))

		argvC = append(argvC, warg)
	}

	argc := C.int(len(argv))
	if updatepath {
		C.PySys_SetArgvEx(argc, (**C.wchar_t)(unsafe.Pointer(&argvC[0])), 1)
	} else {
		C.PySys_SetArgvEx(argc, (**C.wchar_t)(unsafe.Pointer(&argvC[0])), 0)
	}
}

// SetArgv PySys_SetArgv
func SetArgv(argv ...string) {
	if len(argv) <= 0 {
		C.PySys_SetArgv(0, nil)
		return
	}

	argvC := make([]*C.wchar_t, 0, len(argv))
	for _, v := range argv {
		argC := C.CString(v)
		defer C.free(unsafe.Pointer(argC))

		warg := C.Py_DecodeLocale(argC, nil)
		if warg == nil {
			return
		}
		defer C.PyMem_RawFree(unsafe.Pointer(warg))

		argvC = append(argvC, warg)
	}

	argc := C.int(len(argv))
	C.PySys_SetArgv(argc, (**C.wchar_t)(unsafe.Pointer(&argvC[0])))
}
