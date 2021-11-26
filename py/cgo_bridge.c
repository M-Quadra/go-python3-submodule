#include "cgo_bridge.h"

long cgo_Py_REFCNT(PyObject *o) {
    return Py_REFCNT(o);
}