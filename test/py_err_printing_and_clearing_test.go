package main

import (
	"fmt"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pyerr "github.com/M-Quadra/go-python3-submodule/py-err"
	pyexc "github.com/M-Quadra/go-python3-submodule/py-exc"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyErrClear(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyerr.Clear()
	assert.Nil(t, pyerr.Occurred())

	pyerr.SetNone(pyexc.RuntimeError)

	assert.NotNil(t, pyerr.Occurred())
	pyerr.Clear()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrPrintEx(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyerr.Clear()

	pyerr.SetNone(pyexc.RuntimeError)
	pyerr.PrintEx(true)
	pyerr.SetNone(pyexc.RuntimeError)
	pyerr.PrintEx(false)
}

func TestPyErrWriteUnraisable(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyerr.Clear()
	pyerr.WriteUnraisable(nil)

	msgPy := pyunicode.FromString("msg")
	defer py.DecRef(msgPy)

	pyerr.WriteUnraisable(msgPy)
	assert.Nil(t, pyerr.Occurred())
}
