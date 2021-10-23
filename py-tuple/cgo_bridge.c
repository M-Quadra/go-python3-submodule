#include "cgo_bridge.h"

int cgo_PyTuple_Check(PyObject *p) {
    return PyTuple_Check(p);
}

int cgo_PyTuple_CheckExact(PyObject *p) {
    return PyTuple_CheckExact(p);
}