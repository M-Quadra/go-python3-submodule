package pyerr

import (
	/*
		#include "Python.h"
	*/
	"C"
)

// CheckSignals PyErr_CheckSignals
func CheckSignals() int {
	return int(C.PyErr_CheckSignals())
}

// SetInterrupt PyErr_SetInterrupt
func SetInterrupt() {
	C.PyErr_SetInterrupt()
}

// PySignal_SetWakeupFd
