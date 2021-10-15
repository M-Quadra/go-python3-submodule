package pyexc

import (
	/*
		#include "Python.h"
	*/
	"C"
)

var (
	// BaseException PyExc_BaseException
	BaseException = toObject(C.PyExc_BaseException)
	// Exception PyExc_Exception
	Exception = toObject(C.PyExc_Exception)
	// ArithmeticError PyExc_ArithmeticError
	ArithmeticError = toObject(C.PyExc_ArithmeticError)
	// AssertionError PyExc_AssertionError
	AssertionError = toObject(C.PyExc_AssertionError)
	// AttributeError PyExc_AttributeError
	AttributeError = toObject(C.PyExc_AttributeError)
	// BlockingIOError PyExc_BlockingIOError
	BlockingIOError = toObject(C.PyExc_BlockingIOError)
	// BrokenPipeError PyExc_BrokenPipeError
	BrokenPipeError = toObject(C.PyExc_BrokenPipeError)
	// BufferError PyExc_BufferError
	BufferError = toObject(C.PyExc_BufferError)
	// ChildProcessError PyExc_ChildProcessError
	ChildProcessError = toObject(C.PyExc_ChildProcessError)
	// ConnectionAbortedError PyExc_ConnectionAbortedError
	ConnectionAbortedError = toObject(C.PyExc_ConnectionAbortedError)
	// ConnectionError PyExc_ConnectionError
	ConnectionError = toObject(C.PyExc_ConnectionError)
	// ConnectionRefusedError PyExc_ConnectionRefusedError
	ConnectionRefusedError = toObject(C.PyExc_ConnectionRefusedError)
	// ConnectcionResetError PyExc_ConnectionResetError
	ConnectcionResetError = toObject(C.PyExc_ConnectionResetError)
	// EOFError PyExc_EOFError
	EOFError = toObject(C.PyExc_EOFError)
	// FileExistsError PyExc_FileExistsError
	FileExistsError = toObject(C.PyExc_FileExistsError)
	// FileNotFoundError PyExc_FileNotFoundError
	FileNotFoundError = toObject(C.PyExc_FileNotFoundError)
	// FloatingPointError PyExc_FloatingPointError
	FloatingPointError = toObject(C.PyExc_FloatingPointError)
	// GeneratorExit PyExc_GeneratorExit
	GeneratorExit = toObject(C.PyExc_GeneratorExit)
	// ImportError PyExc_ImportError
	ImportError = toObject(C.PyExc_ImportError)
	// IndentationError PyExc_IndentationError
	IndentationError = toObject(C.PyExc_IndentationError)
	// IndexError PyExc_IndexError
	IndexError = toObject(C.PyExc_IndexError)
	// InterruptedError PyExc_InterruptedError
	InterruptedError = toObject(C.PyExc_InterruptedError)
	// IsADirectoryError PyExc_IsADirectoryError
	IsADirectoryError = toObject(C.PyExc_IsADirectoryError)
	// KeyError PyExc_KeyError
	KeyError = toObject(C.PyExc_KeyError)
	// KeyboardInterrupt PyExc_KeyboardInterrupt
	KeyboardInterrupt = toObject(C.PyExc_KeyboardInterrupt)
	// LookupError PyExc_LookupError
	LookupError = toObject(C.PyExc_LookupError)
	// MemoryError PyExc_MemoryError
	MemoryError = toObject(C.PyExc_MemoryError)
	// ModuleNotFoundError PyExc_ModuleNotFoundError
	ModuleNotFoundError = toObject(C.PyExc_ModuleNotFoundError)
	// NameError PyExc_NameError
	NameError = toObject(C.PyExc_NameError)
	// NotADirectoryError PyExc_NotADirectoryError
	NotADirectoryError = toObject(C.PyExc_NotADirectoryError)
	// NotImplementedError PyExc_NotImplementedError
	NotImplementedError = toObject(C.PyExc_NotImplementedError)
	// OSError PyExc_OSError
	OSError = toObject(C.PyExc_OSError)
	// OverflowError PyExc_OverflowError
	OverflowError = toObject(C.PyExc_OverflowError)
	// PermissionError PyExc_PermissionError
	PermissionError = toObject(C.PyExc_PermissionError)
	// ProcessLookupError PyExc_ProcessLookupError
	ProcessLookupError = toObject(C.PyExc_ProcessLookupError)
	// RecursionError PyExc_RecursionError
	RecursionError = toObject(C.PyExc_RecursionError)
	// ReferenceError PyExc_ReferenceError
	ReferenceError = toObject(C.PyExc_ReferenceError)
	// RuntimeError PyExc_RuntimeError
	RuntimeError = toObject(C.PyExc_RuntimeError)
	// StopAsyncIteration PyExc_StopAsyncIteration
	StopAsyncIteration = toObject(C.PyExc_StopAsyncIteration)
	// StopIteration PyExc_StopIteration
	StopIteration = toObject(C.PyExc_StopIteration)
	// SyntaxError PyExc_SyntaxError
	SyntaxError = toObject(C.PyExc_SyntaxError)
	// SystemError PyExc_SystemError
	SystemError = toObject(C.PyExc_SystemError)
	// SystemExit PyExc_SystemExit
	SystemExit = toObject(C.PyExc_SystemExit)
	// TabError PyExc_TabError
	TabError = toObject(C.PyExc_TabError)
	// TimeoutError PyExc_TimeoutError
	TimeoutError = toObject(C.PyExc_TimeoutError)
	// TypeError PyExc_TypeError
	TypeError = toObject(C.PyExc_TypeError)
	// UnboundLocalError PyExc_UnboundLocalError
	UnboundLocalError = toObject(C.PyExc_UnboundLocalError)
	// UnicodeDecodeError PyExc_UnicodeDecodeError
	UnicodeDecodeError = toObject(C.PyExc_UnicodeDecodeError)
	// UnicodeEncodeError PyExc_UnicodeEncodeError
	UnicodeEncodeError = toObject(C.PyExc_UnicodeEncodeError)
	// UnicodeError PyExc_UnicodeError
	UnicodeError = toObject(C.PyExc_UnicodeError)
	// UnicodeTranslateError PyExc_UnicodeTranslateError
	UnicodeTranslateError = toObject(C.PyExc_UnicodeTranslateError)
	// ValueError PyExc_ValueError
	ValueError = toObject(C.PyExc_ValueError)
	// ZeroDivisionError PyExc_ZeroDivisionError
	ZeroDivisionError = toObject(C.PyExc_ZeroDivisionError)
)
