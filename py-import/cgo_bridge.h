#include "Python.h"

PyObject* cgo_PyImport_ImportModuleEx(const char *name, PyObject *globals, PyObject *locals, PyObject *fromlist);