#include "cgo_bridge.h"

int cgo_PyLong_Check(PyObject *p) {
    return PyLong_Check(p);
}

int cgo_PyLong_CheckExact(PyObject *p) {
    return PyLong_CheckExact(p);
}