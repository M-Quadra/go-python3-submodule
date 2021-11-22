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
}

func TestPyBoolFromLong(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	tPy := pybool.FromInt(1)
	defer py.DecRef(tPy)
	assert.Equal(t, py.True, tPy)

	fPy := pybool.FromInt(0)
	defer py.DecRef(fPy)
	assert.Equal(t, py.False, fPy)
}
