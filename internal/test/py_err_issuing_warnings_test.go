package main

import (
	"fmt"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v9/py"
	pydict "github.com/M-Quadra/go-python3-submodule/v9/py-dict"
	pyerr "github.com/M-Quadra/go-python3-submodule/v9/py-err"
	pyexc "github.com/M-Quadra/go-python3-submodule/v9/py-exc"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v9/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyErrWarnEx(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.True(t, pyerr.WarnEx(pyexc.RuntimeWarning, "msg", 1))

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	assert.False(t, pyerr.WarnEx(dic, "msg", 1))
}

func TestPyErrSetImportErrorSubclass(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pyerr.SetImportErrorSubclass(nil, nil, nil, nil))
	pyerr.Clear()

	msg := pyunicode.FromString("no msg")
	defer py.DecRef(msg)

	obj := pyerr.SetImportErrorSubclass(pyexc.ImportError, msg, nil, nil)
	assert.Nil(t, obj)

	assert.NotNil(t, pyerr.Occurred())
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}

func TestPyErrWarnExplicitObject(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	msg := pyunicode.FromString("msg")
	defer py.DecRef(msg)
	msgRefCnt := py.RefCnt(msg)
	defer func() { assert.Equal(t, msgRefCnt, py.RefCnt(msg)) }()

	filename := pyunicode.FromString("test.py")
	defer py.DecRef(filename)
	filenameRefCnt := py.RefCnt(filename)
	defer func() { assert.Equal(t, filenameRefCnt, py.RefCnt(filename)) }()

	module := pyunicode.FromString("model")
	defer py.DecRef(module)
	moduleRefCnt := py.RefCnt(module)
	defer func() { assert.Equal(t, moduleRefCnt, py.RefCnt(module)) }()

	assert.True(t, pyerr.WarnExplicitObject(pyexc.RuntimeWarning, msg, filename, 0, module, nil))

	{ // nil
		assert.True(t, pyerr.WarnExplicitObject(nil, msg, filename, 0, module, nil))

		assert.False(t, pyerr.WarnExplicitObject(nil, nil, filename, 0, module, nil))

		assert.True(t, pyerr.WarnExplicitObject(nil, msg, nil, 0, module, nil))
		assert.True(t, pyerr.WarnExplicitObject(nil, msg, filename, 0, nil, nil))
	}
}

func TestPyErrWarnExplicit(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	msg := pyunicode.FromString("msg")
	defer py.DecRef(msg)
	msgRefCnt := py.RefCnt(msg)
	defer func() { assert.Equal(t, msgRefCnt, py.RefCnt(msg)) }()

	assert.True(t, pyerr.WarnExplicit(pyexc.RuntimeWarning, "msg", "test.py", 1, "module", nil))
	assert.True(t, pyerr.WarnExplicit(nil, "msg", "test.py", 2, "module", nil))
}
