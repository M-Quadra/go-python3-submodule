#include "cgo_bridge.h"

int cgo_PyDict_Check(PyObject *p) {
    return PyDict_Check(p);
}

int cgo_PyDict_CheckExact(PyObject *p) {
    return PyDict_CheckExact(p);
}