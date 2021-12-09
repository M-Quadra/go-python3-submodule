#include "cgo_bridge.h"

long cgo_Py_REFCNT(PyObject *o) {
    return Py_REFCNT(o);
}

int cgo_Py_EnterRecursiveCall(const char *where) {
    return Py_EnterRecursiveCall(where);
}

void cgo_Py_LeaveRecursiveCall() {
    Py_LeaveRecursiveCall();
}