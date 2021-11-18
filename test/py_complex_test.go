package main

import (
	"math"
	"math/rand"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pycallable "github.com/M-Quadra/go-python3-submodule/py-callable"
	pycomplex "github.com/M-Quadra/go-python3-submodule/py-complex"
	pyimport "github.com/M-Quadra/go-python3-submodule/py-import"
	pylist "github.com/M-Quadra/go-python3-submodule/py-list"
	pyobject "github.com/M-Quadra/go-python3-submodule/py-object"
	pytuple "github.com/M-Quadra/go-python3-submodule/py-tuple"
	"github.com/stretchr/testify/assert"
)

func TestPyComplexCheck(t *testing.T) {
	assert.False(t, pycomplex.Check(nil))
	assert.False(t, pycomplex.CheckExact(nil))

	list := pylist.New(0)
	defer py.DecRef(list)
	assert.False(t, pycomplex.Check(list))
	assert.False(t, pycomplex.CheckExact(list))

	v := pycomplex.FromFloat64s(float64(rand.Intn(1000)), float64(rand.Intn(1000)))
	defer py.DecRef(v)

	assert.True(t, pycomplex.Check(v))
	assert.True(t, pycomplex.CheckExact(v))
}

func TestPyComplexFrom(t *testing.T) {
	assert.Zero(t, pycomplex.RealAsFloat64(nil))
	assert.Zero(t, pycomplex.ImagAsFloat64(nil))

	complexTest := pyimport.ImportModule("py_complex_test")
	defer py.DecRef(complexTest)

	funcPy := pyobject.GetAttrString(complexTest, "multiply")
	defer py.DecRef(funcPy)
	assert.True(t, pycallable.Check(funcPy))

	vA := pycomplex.FromFloat64s(0, 1)
	defer py.DecRef(vA)
	vB := pycomplex.FromFloat64s(0, 1)
	defer py.DecRef(vB)

	args := pytuple.FromObjects(vA, vB)
	defer py.DecRef(args)

	result := pyobject.CallObject(funcPy, args)
	defer py.DecRef(result)

	assert.True(t, math.Abs(pycomplex.RealAsFloat64(result)+1) < 1e-8)
}