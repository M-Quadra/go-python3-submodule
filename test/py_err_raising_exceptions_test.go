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

func TestPyErrSetString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pyerr.Occurred())

	pyerr.SetString(nil, "wtf_0")
	assert.Nil(t, pyerr.Occurred())

	pyerr.SetString(pyexc.BaseException, "wtf_1")
	assert.NotNil(t, pyerr.Occurred())

	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrSetObject(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyerr.SetObject(nil, nil)
	assert.Nil(t, pyerr.Occurred())

	msg := pyunicode.FromString("wts")
	defer py.DecRef(msg)

	pyerr.SetObject(pyexc.BaseException, msg)
	assert.NotNil(t, pyerr.Occurred())

	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrSetNone(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pyerr.Occurred())

	pyerr.SetNone(nil)
	assert.Nil(t, pyerr.Occurred())

	{
		msg := pyunicode.FromString("soga")
		defer py.DecRef(msg)
		msgRefCnt := py.RefCnt(msg)
		defer func() { assert.Equal(t, msgRefCnt, py.RefCnt(msg)) }()

		pyerr.SetNone(msg) // SystemError: _PyErr_SetObject: exception 'soga' is not a BaseException subclass

		assert.NotNil(t, pyerr.Occurred())
		pyerr.Print()
		assert.Nil(t, pyerr.Occurred())
	}

	pyerr.SetNone(pyexc.BaseException)
	assert.NotNil(t, pyerr.Occurred())
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrBadArgument(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pyerr.Occurred())

	assert.Equal(t, 0, pyerr.BadArgument())

	assert.NotNil(t, pyerr.Occurred())
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrNoMemory(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pyerr.Occurred())

	pyerr.NoMemory()

	assert.NotNil(t, pyerr.Occurred())
	pyerr.Clear()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrSetImportError(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pyerr.SetImportError(nil, nil, nil))
	pyerr.Clear()

	msg := pyunicode.FromString("msg")
	defer py.DecRef(msg)

	pyerr.SetImportError(msg, nil, nil)
	assert.NotNil(t, pyerr.Occurred())
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrSyntaxLocationObject(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyerr.Clear()

	pyerr.SyntaxLocationObject(nil, 0, 0)
	pyerr.Clear()

	filename := pyunicode.FromString("test.py")
	defer py.DecRef(filename)

	{
		pyerr.SyntaxLocationObject(filename, 0, 0)
		pyerr.Clear()
	}

	pyerr.SetNone(pyexc.SyntaxError)
	pyerr.SyntaxLocationObject(filename, 0, 0)

	assert.NotNil(t, pyerr.Occurred())
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrSyntaxLocationEx(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyerr.Clear()
	pyerr.SyntaxLocationEx("", 0, 0)

	pyerr.SetNone(pyexc.SyntaxError)
	pyerr.SyntaxLocationEx("test.py", 0, 0)

	assert.NotNil(t, pyerr.Occurred())
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrSyntaxLocation(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyerr.Clear()
	pyerr.SyntaxLocation("test.py", 0)

	pyerr.SetNone(pyexc.SyntaxError)
	pyerr.SyntaxLocation("test.py", 0)

	assert.NotNil(t, pyerr.Occurred())
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrBadInternalCall(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyerr.Clear()

	pyerr.BadInternalCall()

	assert.NotNil(t, pyerr.Occurred())
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}
