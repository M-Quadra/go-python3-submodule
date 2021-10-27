#include "Python.h"

int cgo_PyObject_DelAttr(PyObject *o, PyObject *attr_name);

int cgo_PyObject_DelAttrString(PyObject *o, const char *attr_name);

PyObject* cgo_PyObject_CallFunctionObjArgs(PyObject *callable, int argc, PyObject **argv);

PyObject* cgo_PyObject_CallMethodObjArgs(PyObject *callable, PyObject *name, int argc, PyObject **argv);