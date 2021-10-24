package pylist

import (
	python "github.com/M-Quadra/go-python3-submodule"
	pylong "github.com/M-Quadra/go-python3-submodule/py-long"
)

// FromInts []int -> PyList
func FromInts(x []int) *python.PyObject {
	opt := New(len(x))
	for i, v := range x {
		SetItem(opt, i, pylong.FromInt(v))
	}
	return opt
}
