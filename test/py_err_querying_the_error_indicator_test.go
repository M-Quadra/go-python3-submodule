package main

import (
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pyerr "github.com/M-Quadra/go-python3-submodule/py-err"
	pyexc "github.com/M-Quadra/go-python3-submodule/py-exc"
	pyobject "github.com/M-Quadra/go-python3-submodule/py-object"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyErrOccurred(t *testing.T) {
	pyerr.Clear()
	assert.Nil(t, pyerr.Occurred())

	pyerr.SetNone(pyexc.RuntimeError)
	assert.NotNil(t, pyerr.Occurred())
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrExceptionMatches(t *testing.T) {
	pyerr.Clear()
	assert.False(t, pyerr.ExceptionMatches(nil))

	pyerr.SetNone(pyexc.RuntimeError)
	assert.True(t, pyerr.ExceptionMatches(pyexc.RuntimeError))
	pyerr.Clear()
	assert.False(t, pyerr.ExceptionMatches(pyexc.RuntimeError))
}

func TestPyErrGivenExceptionMatches(t *testing.T) {
	assert.True(t, pyerr.GivenExceptionMatches(pyexc.RuntimeError, pyexc.RuntimeError))
	assert.False(t, pyerr.GivenExceptionMatches(pyexc.RuntimeError, pyexc.ArithmeticError))
}

func TestPyErrFetchRestore(t *testing.T) {
	pyerr.Clear()
	exc, value, traceback := pyerr.Fetch()
	assert.Nil(t, exc)
	assert.Nil(t, value)
	assert.Nil(t, traceback)

	pyerr.SetNone(pyexc.RuntimeError)
	assert.NotNil(t, pyerr.Occurred())

	exc, value, traceback = pyerr.Fetch()
	assert.Nil(t, pyerr.Occurred())
	assert.True(t, pyerr.GivenExceptionMatches(pyexc.RuntimeError, exc))
	assert.Nil(t, value)
	assert.Nil(t, traceback)

	pyerr.Restore(exc, value, traceback)
	assert.NotNil(t, pyerr.Occurred())
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrNormalizeException(t *testing.T) {
	pyerr.Clear()

	pyerr.SetNone(pyexc.RuntimeError)
	assert.NotNil(t, pyerr.Occurred())

	exc, value, traceback := pyerr.Fetch()
	assert.True(t, pyerr.GivenExceptionMatches(pyexc.RuntimeError, exc))
	assert.Nil(t, value)
	assert.Nil(t, traceback)

	exc, value, traceback = pyerr.NormalizeException(exc, value, traceback)
	assert.True(t, pyerr.GivenExceptionMatches(pyexc.RuntimeError, exc))
	assert.NotNil(t, value)
	assert.Nil(t, traceback)

	assert.Nil(t, pyerr.Occurred())

	assert.True(t, pyerr.GivenExceptionMatches(exc, pyexc.RuntimeError))
	assert.Equal(t, 1, pyobject.IsInstance(value, exc))
	assert.Nil(t, traceback)

	pyerr.Restore(exc, value, traceback)

	assert.NotNil(t, pyerr.Occurred())
	pyerr.Clear()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrGetSetExcInfo(t *testing.T) {
	pyerr.Clear()

	pyerr.SetNone(pyexc.BufferError)
	assert.NotNil(t, pyerr.Occurred())

	exc, value, traceback := pyerr.GetExcInfo()
	assert.True(t, pyerr.GivenExceptionMatches(exc, py.None), pyunicode.AsString(pyobject.Repr(value)))
	assert.Nil(t, value)
	assert.Nil(t, traceback)

	pyerr.SetExcInfo(exc, value, traceback)
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}
