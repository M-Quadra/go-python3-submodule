#include "cgo_bridge.h"

int cgo_PyByteArray_Check(PyObject *p) {
    return PyByteArray_Check(p);
}

int cgo_PyByteArray_CheckExact(PyObject *p) {
    return PyByteArray_CheckExact(p);
}