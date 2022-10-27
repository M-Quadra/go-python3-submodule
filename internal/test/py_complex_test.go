package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v10/py"
	pycallable "github.com/M-Quadra/go-python3-submodule/v10/py-callable"
	pycomplex "github.com/M-Quadra/go-python3-submodule/v10/py-complex"
	pyimport "github.com/M-Quadra/go-python3-submodule/v10/py-import"
	pylist "github.com/M-Quadra/go-python3-submodule/v10/py-list"
	pyobject "github.com/M-Quadra/go-python3-submodule/v10/py-object"
	pytuple "github.com/M-Quadra/go-python3-submodule/v10/py-tuple"
	"github.com/stretchr/testify/assert"
)

func TestPyComplexCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pycomplex.Check(nil))
	assert.False(t, pycomplex.CheckExact(nil))

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.False(t, pycomplex.Check(list))
	assert.False(t, pycomplex.CheckExact(list))

	v := pycomplex.FromFloat64s(float64(rand.Intn(1000)), float64(rand.Intn(1000)))
	defer py.DecRef(v)
	defer func() { assert.Equal(t, 1, py.RefCnt(v)) }()

	assert.True(t, pycomplex.Check(v))
	assert.True(t, pycomplex.CheckExact(v))
}

func TestPyComplexFrom(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Zero(t, pycomplex.RealAsFloat64(nil))
	assert.Zero(t, pycomplex.ImagAsFloat64(nil))

	complexTest := pyimport.ImportModule("py_complex_test")
	defer py.DecRef(complexTest)
	complexTestRefCnt := py.RefCnt(complexTest)
	defer func() { assert.Equal(t, complexTestRefCnt, py.RefCnt(complexTest)) }()

	funcPy := pyobject.GetAttrString(complexTest, "multiply")
	defer py.DecRef(funcPy)
	funcPyRefCnt := py.RefCnt(funcPy)
	defer func() { assert.Equal(t, funcPyRefCnt, py.RefCnt(funcPy)) }()

	assert.True(t, pycallable.Check(funcPy))

	vA := pycomplex.FromFloat64s(0, 1)
	py.IncRef(vA)
	defer py.DecRef(vA)
	defer func() { assert.Equal(t, 1, py.RefCnt(vA)) }()

	vB := pycomplex.FromFloat64s(0, 1)
	py.IncRef(vB)
	defer py.DecRef(vB)
	defer func() { assert.Equal(t, 1, py.RefCnt(vA)) }()

	args := pytuple.FromObjects(vA, vB)
	defer py.DecRef(args)
	defer func() { assert.Equal(t, 1, py.RefCnt(args)) }()

	result := pyobject.CallObject(funcPy, args)
	defer py.DecRef(result)
	defer func() { assert.Equal(t, 1, py.RefCnt(result)) }()

	assert.True(t, math.Abs(pycomplex.RealAsFloat64(result)+1) < 1e-8)
}
