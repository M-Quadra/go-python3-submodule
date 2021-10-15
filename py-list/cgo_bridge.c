#include "cgo_bridge.h"

int cgo_PyList_Check(PyObject *p) {
    return PyList_Check(p);
}

int cgo_PyList_CheckExact(PyObject *p) {
    return PyList_CheckExact(p);
}