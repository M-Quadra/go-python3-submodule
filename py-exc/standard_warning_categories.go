package pyexc

import (
	/*
		#include "Python.h"
	*/
	"C"
)

var (
	// Warning PyExc_Warning
	Warning = toObject(C.PyExc_Warning)
	// BytesWarning PyExc_BytesWarning
	BytesWarning = toObject(C.PyExc_BytesWarning)
	// DeprecationWarning PyExc_DeprecationWarning
	DeprecationWarning = toObject(C.PyExc_DeprecationWarning)
	// FutureWarning PyExc_FutureWarning
	FutureWarning = toObject(C.PyExc_FutureWarning)
	// ImportWarning PyExc_ImportWarning
	ImportWarning = toObject(C.PyExc_ImportWarning)
	// PendingDeprecationWarning PyExc_PendingDeprecationWarning
	PendingDeprecationWarning = toObject(C.PyExc_PendingDeprecationWarning)
	// ResourceWarning PyExc_ResourceWarning
	ResourceWarning = toObject(C.PyExc_ResourceWarning)
	// RuntimeWarning PyExc_RuntimeWarning
	RuntimeWarning = toObject(C.PyExc_RuntimeWarning)
	// SyntaxWarning PyExc_SyntaxWarning
	SyntaxWarning = toObject(C.PyExc_SyntaxWarning)
	// UnicodeWarning PyExc_UnicodeWarning
	UnicodeWarning = toObject(C.PyExc_UnicodeWarning)
	// UserWarning PyExc_UserWarning
	UserWarning = toObject(C.PyExc_UserWarning)
)
