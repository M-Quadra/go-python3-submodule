#include "cgo_bridge.h"

int cgo_PyBytes_Check(PyObject *p) {
    return PyBytes_Check(p);
}

int cgo_PyBytes_CheckExact(PyObject *p) {
    return PyBytes_CheckExact(p);
}