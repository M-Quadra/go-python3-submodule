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
	fmt.Println(assert.CallerInfo()[0])

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
	fmt.Println(assert.CallerInfo()[0])

	prefix := py.GetPrefix()
	fmt.Println(prefix)
	assert.True(t, len(prefix) > 0)
}

func TestPyGetExecPrefix(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	execPrefix := py.GetExecPrefix()
	fmt.Println(execPrefix)
}

func TestPyGetProgramFullPath(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	fullPath := py.GetProgramFullPath()
	fmt.Println(fullPath)
	assert.True(t, len(fullPath) > 0)
}

func TestPyGetSetPath(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	defaultPath := py.GetPath()
	defer py.SetPath(defaultPath)

	name := "påth"
	py.SetPath(name)
	assert.Equal(t, name, py.GetPath())
}

func TestPyGetVersion(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	version := py.GetVersion()
	fmt.Println(version)
	assert.True(t, len(version) > 0)
}

func TestPyGetPlatform(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	platform := py.GetPlatform()
	fmt.Println(platform)
	assert.True(t, len(platform) > 0)
}

func TestPyGetCopyright(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	copyright := py.GetCopyright()
	fmt.Println(copyright)
	assert.True(t, len(copyright) > 0)
}

func TestPyGetCompiler(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	compiler := py.GetCompiler()
	fmt.Println(compiler)
	assert.True(t, len(compiler) > 0)
}

func TestPyGetBuildInfo(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	buildInfo := py.GetBuildInfo()
	fmt.Println(buildInfo)
	assert.True(t, len(buildInfo) > 0)
}

func TestPySysSetArgvEx(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	{ // argv recovery
		path := pysys.GetObject("path")
		py.IncRef(path)
		defer py.DecRef(path)
		defer func() { assert.Equal(t, 1, py.RefCnt(path)) }()

		oldPath := pylist.New(pylist.Size(path))
		defer func() { assert.Equal(t, 1, py.RefCnt(oldPath)) }()

		assert.True(t, pylist.SetSlice(oldPath, 0, pylist.Size(path), path))
		defer func() {
			assert.Equal(t, 1, py.RefCnt(oldPath))
			assert.True(t, pysys.SetObject("path", oldPath))
			py.DecRef(oldPath)
		}()
	}

	pysys.SetArgvEx(false, "test.py")

	argv := pysys.GetObject("argv")
	defer func() { assert.Equal(t, 1, py.RefCnt(argv)) }()

	assert.Equal(t, 1, pylist.Size(argv))
	assert.Equal(t, "test.py", pyunicode.AsString(pylist.GetItem(argv, 0)))
}

func TestPySysSetArgv(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	{ // path recovery
		path := pysys.GetObject("path")
		py.IncRef(path)
		defer py.DecRef(path)
		defer func() { assert.Equal(t, 1, py.RefCnt(path)) }()

		oldPath := pylist.New(pylist.Size(path))
		defer func() { assert.Equal(t, 1, py.RefCnt(oldPath)) }()

		assert.True(t, pylist.SetSlice(oldPath, 0, pylist.Size(path), path))
		defer func() {
			assert.Equal(t, 1, py.RefCnt(oldPath))
			assert.True(t, pysys.SetObject("path", oldPath))
			py.DecRef(oldPath)
		}()
	}

	{ // argv recovery
		argv := pysys.GetObject("argv")
		py.IncRef(argv)
		defer py.DecRef(argv)

		argvArr := make([]string, 0, pylist.Size(argv))
		for i := 0; i < pylist.Size(argv); i++ {
			argvArr = append(argvArr, pyunicode.AsString(pylist.GetItem(argv, i)))
		}
		defer func() {
			pysys.SetArgvEx(false, argvArr...)
			assert.Equal(t, 1, py.RefCnt(argv))
		}()
	}

	pysys.SetArgv("test.py")

	argv := pysys.GetObject("argv")
	defer func() { assert.Equal(t, 1, py.RefCnt(argv)) }()

	assert.Equal(t, 1, pylist.Size(argv))
	assert.Equal(t, "test.py", pyunicode.AsString(pylist.GetItem(argv, 0)))
}

func TestPyGetSetPythonHome(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	defaultHome := py.GetPythonHome()
	defer py.SetPythonHome(defaultHome)

	home := "høme"
	py.SetPythonHome(home)
	assert.Equal(t, home, py.GetPythonHome())
}
