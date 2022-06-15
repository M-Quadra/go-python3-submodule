package main

import (
	"fmt"
	"testing"

	pycallable "github.com/M-Quadra/go-python3-submodule/v9/py-callable"
	pydict "github.com/M-Quadra/go-python3-submodule/v9/py-dict"
	pyeval "github.com/M-Quadra/go-python3-submodule/v9/py-eval"
	"github.com/stretchr/testify/assert"
)

func TestPyEvalGetBuiltins(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	builtinsPy := pyeval.GetBuiltins()
	assert.NotNil(t, builtinsPy)

	lenPy := pydict.GetItemString(builtinsPy, "len")
	assert.True(t, pycallable.Check(lenPy))
}

func TestPyEvalGetLocals(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	localsPy := pyeval.GetLocals()
	assert.Nil(t, localsPy)
}

func TestPyEvalGetGlobals(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	globalsPy := pyeval.GetGlobals()
	assert.Nil(t, globalsPy)
}

func TestPyEvalGetFuncName(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	builtinsPy := pyeval.GetBuiltins()
	assert.NotNil(t, builtinsPy)

	lenPy := pydict.GetItemString(builtinsPy, "len")
	assert.True(t, pycallable.Check(lenPy))
	assert.Equal(t, "len", pyeval.GetFuncName(lenPy))

	assert.Equal(t, "", pyeval.GetFuncName(nil))
}

func TestPyEvalGetFuncDesc(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	builtinsPy := pyeval.GetBuiltins()
	assert.NotNil(t, builtinsPy)

	lenPy := pydict.GetItemString(builtinsPy, "len")
	assert.True(t, pycallable.Check(lenPy))

	assert.Equal(t, "()", pyeval.GetFuncDesc(lenPy))
}
