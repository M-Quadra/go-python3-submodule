#include "cgo_bridge.h"

int cgo_PyFloat_Check(PyObject *p) {
    return PyFloat_Check(p);
}

int cgo_PyFloat_CheckExact(PyObject *p) {
    return PyFloat_CheckExact(p);
}