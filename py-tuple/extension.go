package pytuple

import python "github.com/M-Quadra/go-python3-submodule/v8"

// FromObjects ...*python.PyObject -> *PyTuple
func FromObjects(objs ...*python.PyObject) *python.PyObject {
	tuple := New(len(objs))
	for i := range objs {
		SetItem(tuple, i, objs[i])
	}
	return tuple
}
