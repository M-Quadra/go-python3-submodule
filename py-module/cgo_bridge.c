#include "cgo_bridge.h"

int cgo_PyModule_Check(PyObject *p) {
    return PyModule_Check(p);
}

int cgo_PyModule_CheckExact(PyObject *p) {
    return PyModule_CheckExact(p);
}
