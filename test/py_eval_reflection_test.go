package main

import (
	"testing"

	pycallable "github.com/M-Quadra/go-python3-submodule/py-callable"
	pydict "github.com/M-Quadra/go-python3-submodule/py-dict"
	pyeval "github.com/M-Quadra/go-python3-submodule/py-eval"
	"github.com/stretchr/testify/assert"
)

func TestPyEvalGetBuiltins(t *testing.T) {
	builtinsPy := pyeval.GetBuiltins()
	assert.NotNil(t, builtinsPy)

	lenPy := pydict.GetItemString(builtinsPy, "len")
	assert.True(t, pycallable.Check(lenPy))
}

func TestPyEvalGetLocals(t *testing.T) {
	localsPy := pyeval.GetLocals()
	assert.Nil(t, localsPy)
}

func TestPyEvalGetGlobals(t *testing.T) {
	globalsPy := pyeval.GetGlobals()
	assert.Nil(t, globalsPy)
}

func TestPyEvalGetFuncName(t *testing.T) {
	builtinsPy := pyeval.GetBuiltins()
	assert.NotNil(t, builtinsPy)

	lenPy := pydict.GetItemString(builtinsPy, "len")
	assert.True(t, pycallable.Check(lenPy))
	assert.Equal(t, "len", pyeval.GetFuncName(lenPy))

	assert.Equal(t, "", pyeval.GetFuncName(nil))
}

func TestPyEvalGetFuncDesc(t *testing.T) {
	builtinsPy := pyeval.GetBuiltins()
	assert.NotNil(t, builtinsPy)

	lenPy := pydict.GetItemString(builtinsPy, "len")
	assert.True(t, pycallable.Check(lenPy))

	assert.Equal(t, "()", pyeval.GetFuncDesc(lenPy))
}
