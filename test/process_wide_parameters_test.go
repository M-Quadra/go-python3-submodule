package main

import (
	"fmt"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pylist "github.com/M-Quadra/go-python3-submodule/py-list"
	pysys "github.com/M-Quadra/go-python3-submodule/py-sys"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyGetSetProgramName(t *testing.T) {
	defaultName := py.GetProgramName()
	assert.Equal(t, "python3", defaultName)
	defer py.SetProgramName(defaultName)

	name := "py3∑åß∂"
	py.SetProgramName(name)
	assert.Equal(t, name, py.GetProgramName())

	py.SetProgramName("")
	assert.Equal(t, name, py.GetProgramName())
}

func TestPyGetPrefix(t *testing.T) {
	prefix := py.GetPrefix()
	fmt.Println(prefix)
	assert.True(t, len(prefix) > 0)
}

func TestPyGetExecPrefix(t *testing.T) {
	execPrefix := py.GetExecPrefix()
	fmt.Println(execPrefix)
}

func TestPyGetProgramFullPath(t *testing.T) {
	fullPath := py.GetProgramFullPath()
	fmt.Println(fullPath)
	assert.True(t, len(fullPath) > 0)
}

func TestPyGetSetPath(t *testing.T) {
	defaultPath := py.GetPath()
	fmt.Println(defaultPath)
	defer py.SetPath(defaultPath)

	name := "påth"
	py.SetPath(name)
	assert.Equal(t, name, py.GetPath())
}

func TestPyGetVersion(t *testing.T) {
	version := py.GetVersion()
	fmt.Println(version)
	assert.True(t, len(version) > 0)
}

func TestPyGetPlatform(t *testing.T) {
	platform := py.GetPlatform()
	fmt.Println(platform)
	assert.True(t, len(platform) > 0)
}

func TestPyGetCopyright(t *testing.T) {
	copyright := py.GetCopyright()
	fmt.Println(copyright)
	assert.True(t, len(copyright) > 0)
}

func TestPyGetCompiler(t *testing.T) {
	compiler := py.GetCompiler()
	fmt.Println(compiler)
	assert.True(t, len(compiler) > 0)
}

func TestPyGetBuildInfo(t *testing.T) {
	buildInfo := py.GetBuildInfo()
	fmt.Println(buildInfo)
	assert.True(t, len(buildInfo) > 0)
}

func TestPySysSetArgvEx(t *testing.T) {
	pysys.SetArgvEx(false, "test.py")

	argv := pysys.GetObject("argv")
	assert.Equal(t, 1, pylist.Size(argv))
	assert.Equal(t, "test.py", pyunicode.AsString(pylist.GetItem(argv, 0)))
}

func TestPySysSetArgv(t *testing.T) {
	pysys.SetArgv("test.py")

	argv := pysys.GetObject("argv")
	assert.Equal(t, 1, pylist.Size(argv))
	assert.Equal(t, "test.py", pyunicode.AsString(pylist.GetItem(argv, 0)))
}

func TestPyGetSetPythonHome(t *testing.T) {
	defaultHome := py.GetPythonHome()
	defer py.SetPythonHome(defaultHome)

	home := "høme"
	py.SetPythonHome(home)
	assert.Equal(t, home, py.GetPythonHome())
}
