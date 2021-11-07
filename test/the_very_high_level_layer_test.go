package main

import (
	"embed"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pyobject "github.com/M-Quadra/go-python3-submodule/py-object"
	pyrun "github.com/M-Quadra/go-python3-submodule/py-run"
	pysys "github.com/M-Quadra/go-python3-submodule/py-sys"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

// func TestPyMain(t *testing.T) {
// 	fmt.Println("???")
// 	py.Main("test.py")
// }

func TestPyfunc(t *testing.T) {
	exitCode := pyrun.AnyFile("test.py")
	assert.Equal(t, 0, exitCode)

	stdout := pysys.GetObject("stdout")

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
	exitCode := pyrun.SimpleString(_testPy)
	assert.Equal(t, 0, exitCode)

	stdout := pysys.GetObject("stdout")

	funName := pyunicode.FromString("getvalue")
	defer py.DecRef(funName)

	result := pyobject.CallMethodNoArgs(stdout, funName)
	defer py.DecRef(result)

	assert.Equal(t, "hello world\n", pyunicode.AsString(result))
}
