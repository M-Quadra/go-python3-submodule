package py

import (
	/*
		#include "Python.h"
	*/
	"C"
)

// Main Py_Main
// func Main(argv ...string) int {
// 	argc := C.int(len(argv))
// 	argvC := make([]*C.wchar_t, 0, argc)
// 	for _, v := range argv {
// 		strC := C.CString(v)
// 		warg := C.Py_DecodeLocale(strC, nil)
// 		C.free(unsafe.Pointer(strC))
// 		if warg == nil {
// 			return -1
// 		}
// 		defer C.PyMem_RawFree(unsafe.Pointer(warg))

// 		argvC = append(argvC, warg)
// 	}

// 	// can't run, why ????
// 	return int(C.Py_Main(argc, (**C.wchar_t)(unsafe.Pointer(&argvC[0]))))
// }

// Py_BytesMain
