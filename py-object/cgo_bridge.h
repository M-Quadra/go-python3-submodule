#include "Python.h"

int cgo_PyObject_DelAttr(PyObject *o, PyObject *attr_name);

int cgo_PyObject_DelAttrString(PyObject *o, const char *attr_name);
