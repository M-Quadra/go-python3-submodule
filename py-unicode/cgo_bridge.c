#include "cgo_bridge.h"

int cgo_PyUnicode_Check(PyObject *o) {
    return PyUnicode_Check(o);
}

int cgo_PyUnicode_CheckExact(PyObject *o) {
    return PyUnicode_CheckExact(o);
}
