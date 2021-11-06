package py

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import "unsafe"

// [DataDog/go-python3]: The argument for Py_SetProgramName, Py_SetPath and Py_SetPythonHome should point to a zero-terminated wide character string in static storage
// [DataDog/go-python3]: whose contents will not change for the duration of the program’s execution
var (
	programName *C.wchar_t
	pythonPath  *C.wchar_t
	pythonHome  *C.wchar_t
)

// SetStandardStreamEncoding Py_SetStandardStreamEncoding
func SetStandardStreamEncoding(encoding, errors string) int {
	encodingC := C.CString(encoding)
	defer C.free(unsafe.Pointer(encodingC))

	errorsC := C.CString(errors)
	defer C.free(unsafe.Pointer(errorsC))

	return int(C.Py_SetStandardStreamEncoding(encodingC, errorsC))
}

// SetProgramName Py_SetProgramName
func SetProgramName(name string) {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	newProgramName := C.Py_DecodeLocale(nameC, nil)
	if newProgramName == nil {
		return
	}
	C.Py_SetProgramName(newProgramName)

	// [DataDog/go-python3]: no operation is performed if nil
	C.PyMem_RawFree(unsafe.Pointer(programName))
	programName = newProgramName
	return
}

// GetProgramName Py_GetProgramName
func GetProgramName() string {
	wcname := C.Py_GetProgramName()
	if wcname == nil {
		return ""
	}

	nameC := C.Py_EncodeLocale(wcname, nil)
	if nameC == nil {
		return ""
	}
	defer C.PyMem_Free(unsafe.Pointer(nameC))

	return C.GoString(nameC)
}

// GetPrefix Py_GetPrefix
func GetPrefix() string {
	wcname := C.Py_GetPrefix()
	if wcname == nil {
		return ""
	}

	nameC := C.Py_EncodeLocale(wcname, nil)
	if nameC == nil {
		return ""
	}
	defer C.PyMem_Free(unsafe.Pointer(nameC))

	return C.GoString(nameC)
}

// GetExecPrefix Py_GetExecPrefix
func GetExecPrefix() string {
	wcname := C.Py_GetExecPrefix()
	if wcname == nil {
		return ""
	}

	nameC := C.Py_EncodeLocale(wcname, nil)
	if nameC == nil {
		return ""
	}
	defer C.PyMem_Free(unsafe.Pointer(nameC))

	return C.GoString(nameC)
}

// GetProgramFullPath Py_GetProgramFullPath
func GetProgramFullPath() string {
	wcname := C.Py_GetProgramFullPath()
	if wcname == nil {
		return ""
	}

	nameC := C.Py_EncodeLocale(wcname, nil)
	if nameC == nil {
		return ""
	}
	defer C.PyMem_Free(unsafe.Pointer(nameC))

	return C.GoString(nameC)
}

// GetPath Py_GetPath
func GetPath() string {
	wcname := C.Py_GetPath()
	if wcname == nil {
		return ""
	}

	nameC := C.Py_EncodeLocale(wcname, nil)
	if nameC == nil {
		return ""
	}
	defer C.PyMem_Free(unsafe.Pointer(nameC))

	return C.GoString(nameC)
}

// SetPath Py_SetPath
func SetPath(path string) {
	pathC := C.CString(path)
	defer C.free(unsafe.Pointer(pathC))

	newPath := C.Py_DecodeLocale(pathC, nil)
	if newPath == nil {
		return
	}
	C.Py_SetPath(newPath)

	C.PyMem_RawFree(unsafe.Pointer(pythonPath))
	pythonHome = newPath
}

// GetVersion Py_GetVersion
func GetVersion() string {
	versionC := C.Py_GetVersion()
	return C.GoString(versionC)
}

// GetPlatform Py_GetPlatform
func GetPlatform() string {
	platformC := C.Py_GetPlatform()
	return C.GoString(platformC)
}

// GetCopyright Py_GetCopyright
func GetCopyright() string {
	copyrightC := C.Py_GetCopyright()
	return C.GoString(copyrightC)
}

// GetCompiler Py_GetCompiler
func GetCompiler() string {
	compilerC := C.Py_GetCompiler()
	return C.GoString(compilerC)
}

// GetBuildInfo Py_GetBuildInfo
func GetBuildInfo() string {
	buildInfoC := C.Py_GetBuildInfo()
	return C.GoString(buildInfoC)
}

// SetPythonHome Py_SetPythonHome
func SetPythonHome(home string) {
	homeC := C.CString(home)
	defer C.free(unsafe.Pointer(homeC))

	newHome := C.Py_DecodeLocale(homeC, nil)
	if newHome == nil {
		return
	}
	C.Py_SetPythonHome(newHome)

	C.PyMem_RawFree(unsafe.Pointer(pythonHome))
	pythonHome = newHome
}

// GetPythonHome Py_GetPythonHome
func GetPythonHome() string {
	wchome := C.Py_GetPythonHome()
	if wchome == nil {
		return ""
	}

	homeC := C.Py_EncodeLocale(wchome, nil)
	if homeC == nil {
		return ""
	}
	defer C.PyMem_Free(unsafe.Pointer(homeC))

	return C.GoString(homeC)
}
