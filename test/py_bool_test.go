package main

import (
	"fmt"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pybool "github.com/M-Quadra/go-python3-submodule/py-bool"
	pylist "github.com/M-Quadra/go-python3-submodule/py-list"
	"github.com/stretchr/testify/assert"
)

func TestPyBoolCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pybool.Check(nil))

	list := pylist.New(1)
	defer py.DecRef(list)
	assert.True(t, pybool.Check(list)) // ???

	assert.Equal(t, 1, py.RefCnt(list))
}

func TestPyBoolFromLong(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	tPy := pybool.FromInt(1)
	defer py.DecRef(tPy)
	refCntTrue := py.RefCnt(py.True)
	defer func() { assert.Equal(t, refCntTrue, py.RefCnt(py.True)) }()

	assert.Equal(t, py.True, tPy)

	fPy := pybool.FromInt(0)
	defer py.DecRef(fPy)
	refCntFalse := py.RefCnt(py.False)
	defer func() { assert.Equal(t, refCntFalse, py.RefCnt(py.False)) }()

	assert.Equal(t, py.False, fPy)
}
