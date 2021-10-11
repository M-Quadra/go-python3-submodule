#include "cgo_bridge.h"

PyObject* cgo_PyImport_ImportModuleEx(const char *name, PyObject *globals, PyObject *locals, PyObject *fromlist) {
    return PyImport_ImportModuleEx(name, globals, locals, fromlist);
}