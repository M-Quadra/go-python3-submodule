package main

import (
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pydict "github.com/M-Quadra/go-python3-submodule/py-dict"
	pyerr "github.com/M-Quadra/go-python3-submodule/py-err"
	pyexc "github.com/M-Quadra/go-python3-submodule/py-exc"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyErrWarnEx(t *testing.T) {
	assert.True(t, pyerr.WarnEx(pyexc.RuntimeWarning, "msg", 1))

	dic := pydict.New()
	defer py.DecRef(dic)
	assert.False(t, pyerr.WarnEx(dic, "msg", 1))
}

func TestPyErrSetImportErrorSubclass(t *testing.T) {
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
	msg := pyunicode.FromString("msg")
	defer py.DecRef(msg)

	filename := pyunicode.FromString("test.py")
	defer py.DecRef(filename)

	module := pyunicode.FromString("model")
	defer py.DecRef(module)

	assert.True(t, pyerr.WarnExplicitObject(pyexc.RuntimeWarning, msg, filename, 0, module, nil))

	{ // nil
		assert.True(t, pyerr.WarnExplicitObject(nil, msg, filename, 0, module, nil))

		assert.False(t, pyerr.WarnExplicitObject(nil, nil, filename, 0, module, nil))

		assert.True(t, pyerr.WarnExplicitObject(nil, msg, nil, 0, module, nil))
		assert.True(t, pyerr.WarnExplicitObject(nil, msg, filename, 0, nil, nil))
	}
}

func TestPyErrWarnExplicit(t *testing.T) {
	msg := pyunicode.FromString("msg")
	defer py.DecRef(msg)

	assert.True(t, pyerr.WarnExplicit(pyexc.RuntimeWarning, "msg", "test.py", 1, "module", nil))
	assert.True(t, pyerr.WarnExplicit(nil, "msg", "test.py", 2, "module", nil))
}
