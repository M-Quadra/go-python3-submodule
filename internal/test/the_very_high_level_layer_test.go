package main

import (
	"embed"
	"fmt"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v11/py"
	pyobject "github.com/M-Quadra/go-python3-submodule/v11/py-object"
	pyrun "github.com/M-Quadra/go-python3-submodule/v11/py-run"
	pysys "github.com/M-Quadra/go-python3-submodule/v11/py-sys"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v11/py-unicode"
	"github.com/stretchr/testify/assert"
)

// func TestPyMain(t *testing.T) {
// 	fmt.Println("current:", assert.CallerInfo()[0])

// 	fmt.Println("???")
// 	py.Main("test.py")
// }

func TestPyfunc(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	exitCode, err := pyrun.AnyFile("test_233.py")
	assert.NotNil(t, err)
	assert.Equal(t, -1, exitCode)

	exitCode, err = pyrun.AnyFile("test.py")
	assert.Nil(t, err)
	assert.Equal(t, 0, exitCode)

	stdout := pysys.GetObject("stdout")
	stdoutRefCnt := py.RefCnt(stdout)
	defer func() { assert.Equal(t, stdoutRefCnt, py.RefCnt(stdout)) }()

	funName := pyunicode.FromString("getvalue")
	defer py.DecRef(funName)

	result := pyobject.CallMethodNoArgs(stdout, funName)
	defer py.DecRef(result)

	assert.Equal(t, "hello world\n", pyunicode.AsString(result))
}

var (
	_ embed.FS
	//go:embed test.py
	// test.py
	_testPy string
)

func TestPyRunSimpleString(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	exitCode := pyrun.SimpleString(_testPy)
	assert.Equal(t, 0, exitCode)

	stdout := pysys.GetObject("stdout")
	stdoutRefCnt := py.RefCnt(stdout)
	defer func() { assert.Equal(t, stdoutRefCnt, py.RefCnt(stdout)) }()

	funName := pyunicode.FromString("getvalue")
	defer py.DecRef(funName)

	result := pyobject.CallMethodNoArgs(stdout, funName)
	defer py.DecRef(result)

	assert.Equal(t, "hello world\n", pyunicode.AsString(result))
}
