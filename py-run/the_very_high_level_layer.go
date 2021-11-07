package pyrun

import (
	/*
		#include "Python.h"
	*/
	"C"
)

import "unsafe"

// AnyFile PyRun_AnyFile
func AnyFile(filename string) int {
	filenameC := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))

	mode := C.CString("r")
	defer C.free(unsafe.Pointer(mode))

	fileC, err := C.fopen(filenameC, mode)
	if err != nil {
		return -1
	}
	defer C.fclose(fileC)

	return int(C.PyRun_AnyFile(fileC, filenameC))
}

// PyRun_AnyFileFlags

// PyRun_AnyFileEx

// PyRun_AnyFileExFlags

// SimpleString PyRun_SimpleString
func SimpleString(command string) int {
	commandC := C.CString(command)
	defer C.free(unsafe.Pointer(commandC))

	return int(C.PyRun_SimpleString(commandC))
}

// PyRun_SimpleStringFlags

// PyRun_SimpleFile

// PyRun_SimpleFileEx

// PyRun_SimpleFileExFlags

// PyRun_InteractiveOne

// PyRun_InteractiveOneFlags

// PyRun_InteractiveLoop

// PyRun_InteractiveLoopFlags

// PyRun_String

// PyRun_StringFlags

// PyRun_File

// PyRun_FileEx

// PyRun_FileFlags

// PyRun_FileExFlags
