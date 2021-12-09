#include "Python.h"

long cgo_Py_REFCNT(PyObject *o);

int cgo_Py_EnterRecursiveCall(const char *where);

void cgo_Py_LeaveRecursiveCall();