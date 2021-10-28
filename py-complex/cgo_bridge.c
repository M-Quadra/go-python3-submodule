#include "cgo_bridge.h"

int cgo_PyComplex_Check(PyObject *p) {
    return PyComplex_Check(p);
}

int cgo_PyComplex_CheckExact(PyObject *p) {
    return PyComplex_CheckExact(p);
}