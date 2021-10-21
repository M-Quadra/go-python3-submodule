#include "cgo_bridge.h"

int cgo_PyBool_Check(PyObject *p) {
    return PyBool_Check(p);
}
