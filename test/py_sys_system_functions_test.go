package main

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/M-Quadra/go-python3-submodule/py"
	pydict "github.com/M-Quadra/go-python3-submodule/py-dict"
	pylist "github.com/M-Quadra/go-python3-submodule/py-list"
	pysys "github.com/M-Quadra/go-python3-submodule/py-sys"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPySysGetSetObject(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	nPlatform := pyunicode.FromString("test")
	defer py.DecRef(nPlatform)
	nPlatformRefCnt := py.RefCnt(nPlatform)
	defer func() { assert.Equal(t, nPlatformRefCnt, py.RefCnt(nPlatform)) }()

	name := "platform"
	platform := pysys.GetObject(name)
	platformRefCnt := py.RefCnt(platform)
	defer func() { assert.Equal(t, platformRefCnt, py.RefCnt(platform)) }()
	assert.True(t, pyunicode.Check(platform))

	assert.True(t, pysys.SetObject(name, nil))
	assert.Nil(t, pysys.GetObject(name))
	assert.True(t, pysys.SetObject(name, nil))

	assert.True(t, pysys.SetObject(name, nPlatform))
	assert.Equal(t, nPlatform, pysys.GetObject(name))
	assert.True(t, pysys.SetObject(name, platform))
}

func TestPySysWarnOptions(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	str := "ignore"
	strPy := pyunicode.FromString(str)
	defer py.DecRef(strPy)
	strPyRefCnt := py.RefCnt(strPy)
	defer func() { assert.Equal(t, strPyRefCnt, py.RefCnt(strPy)) }()

	warnoptions := pysys.GetObject("warnoptions")
	assert.Zero(t, pylist.Size(warnoptions))

	pysys.AddWarnOption(str)
	warnoptions = pysys.GetObject("warnoptions")
	assert.Equal(t, str, pyunicode.AsString(pylist.GetItem(warnoptions, 0)))

	pysys.ResetWarnOptions()
	warnoptions = pysys.GetObject("warnoptions")
	assert.Zero(t, pylist.Size(warnoptions))

	pysys.AddWarnOptionUnicode(nil)
	pysys.AddWarnOptionUnicode(strPy)
	warnoptions = pysys.GetObject("warnoptions")
	assert.Equal(t, str, pyunicode.AsString(pylist.GetItem(warnoptions, 0)))

	pysys.ResetWarnOptions()
	warnoptions = pysys.GetObject("warnoptions")
	assert.Zero(t, pylist.Size(warnoptions))
}

func TestPySysSetPath(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	path := pysys.GetObject("path")
	pathRefCnt := py.RefCnt(path)
	defer func() { assert.Equal(t, pathRefCnt, py.RefCnt(path)) }()
	py.IncRef(path)
	defer py.DecRef(path)

	pysys.SetPath("test")
	nPath := pysys.GetObject("path")
	assert.Equal(t, "test", pyunicode.AsString(pylist.GetItem(nPath, 0)))
	assert.True(t, unsafe.Pointer(path) != unsafe.Pointer(nPath))

	pysys.SetObject("path", path)
}

func TestPySysXOption(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	xOptions := pysys.GetXOptions()
	xOptionsRefCnt := py.RefCnt(xOptions)
	defer func() { assert.Equal(t, xOptionsRefCnt, py.RefCnt(xOptions)) }()
	size := pydict.Size(xOptions)
	defer func() { assert.Equal(t, size, pydict.Size(xOptions)) }()

	pysys.AddXOption("faulthandler")
	faulthandler := pydict.GetItemString(xOptions, "faulthandler")
	assert.Equal(t, py.True, faulthandler)
	pydict.Clear(xOptions)
}
