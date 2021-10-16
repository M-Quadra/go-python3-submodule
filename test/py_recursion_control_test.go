package main

import (
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyRecursiveCall(t *testing.T) {
	assert.Zero(t, py.EnterRecursiveCall("in test function"))
	py.LeaveRecursiveCall()
}

func TestPyRepr(t *testing.T) {
	strPy := pyunicode.FromString("hello world")
	defer py.DecRef(strPy)

	assert.Zero(t, py.ReprEnter(strPy))
	assert.True(t, py.ReprEnter(strPy) > 0)

	py.ReprLeave(strPy)
	py.ReprLeave(strPy)
}
