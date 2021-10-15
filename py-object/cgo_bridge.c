#include "cgo_bridge.h"

int cgo_PyObject_DelAttr(PyObject *o, PyObject *attr_name) {
    return PyObject_DelAttr(o, attr_name);
}

int cgo_PyObject_DelAttrString(PyObject *o, const char *attr_name) {
    return PyObject_DelAttrString(o, attr_name);
}