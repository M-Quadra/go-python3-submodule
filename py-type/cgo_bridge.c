#include "cgo_bridge.h"

int cgo_PyType_Check(PyObject *o) {
    return PyType_Check(o);
}

int cgo_PyType_CheckExact(PyObject *o) {
    return PyType_CheckExact(o);
}