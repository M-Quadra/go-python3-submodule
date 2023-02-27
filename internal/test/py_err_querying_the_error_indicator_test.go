package main

import (
	"fmt"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v11/py"
	pyerr "github.com/M-Quadra/go-python3-submodule/v11/py-err"
	pyexc "github.com/M-Quadra/go-python3-submodule/v11/py-exc"
	pyobject "github.com/M-Quadra/go-python3-submodule/v11/py-object"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v11/py-unicode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPyErrOccurred(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	pyerr.Clear()
	assert.Nil(t, pyerr.Occurred())

	pyerr.SetNone(pyexc.RuntimeError)
	assert.NotNil(t, pyerr.Occurred())
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrExceptionMatches(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	pyerr.Clear()
	assert.False(t, pyerr.ExceptionMatches(nil))

	pyerr.SetNone(pyexc.RuntimeError)
	assert.True(t, pyerr.ExceptionMatches(pyexc.RuntimeError))
	pyerr.Clear()
	assert.False(t, pyerr.ExceptionMatches(pyexc.RuntimeError))
}

func TestPyErrGivenExceptionMatches(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	assert.True(t, pyerr.GivenExceptionMatches(pyexc.RuntimeError, pyexc.RuntimeError))
	assert.False(t, pyerr.GivenExceptionMatches(pyexc.RuntimeError, pyexc.ArithmeticError))
}

func TestPyErrFetchRestore(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

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
	fmt.Println("current:", assert.CallerInfo()[0])

	pyerr.Clear()

	pyerr.SetNone(pyexc.RuntimeError)
	require.NotNil(t, pyerr.Occurred())

	exc, value, traceback := pyerr.Fetch()
	require.True(t, pyerr.GivenExceptionMatches(pyexc.RuntimeError, exc))
	require.Nil(t, value)
	require.Nil(t, traceback)

	exc, value, traceback = pyerr.NormalizeException(exc, value, traceback)
	require.True(t, pyerr.GivenExceptionMatches(pyexc.RuntimeError, exc))
	require.NotNil(t, value)
	require.Nil(t, traceback)

	require.Nil(t, pyerr.Occurred())

	require.True(t, pyerr.GivenExceptionMatches(exc, pyexc.RuntimeError))
	require.Equal(t, 1, pyobject.IsInstance(value, exc))
	require.Nil(t, traceback)

	pyerr.Restore(exc, value, traceback)

	require.NotNil(t, pyerr.Occurred())
	pyerr.Clear()
	require.Nil(t, pyerr.Occurred())
}

func TestPyErrGetSetExcInfo(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	pyerr.Clear()

	pyerr.SetNone(pyexc.BufferError)
	require.NotNil(t, pyerr.Occurred())

	exc, value, traceback := pyerr.GetExcInfo()
	assert.True(t, pyerr.GivenExceptionMatches(exc, py.None), pyunicode.AsString(pyobject.Repr(value)))
	assert.Nil(t, value)
	assert.Nil(t, traceback)

	pyerr.SetExcInfo(exc, value, traceback)
	pyerr.Print()
	require.Nil(t, pyerr.Occurred())
}
